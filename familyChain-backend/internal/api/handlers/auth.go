package handlers

import (
	"net/http"
	"strings"

	"eth-for-babies-backend/internal/models"
	"eth-for-babies-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
	db         *gorm.DB
	jwtManager *utils.JWTManager
}

func NewAuthHandler(db *gorm.DB, jwtManager *utils.JWTManager) *AuthHandler {
	return &AuthHandler{
		db:         db,
		jwtManager: jwtManager,
	}
}

type GetNonceRequest struct {
	WalletAddress string `uri:"wallet_address" binding:"required"`
}

type LoginRequest struct {
	WalletAddress string `json:"wallet_address" binding:"required"`
	Signature     string `json:"signature" binding:"required"`
	Role          string `json:"role,omitempty"`
}

type RegisterRequest struct {
	WalletAddress string `json:"wallet_address" binding:"required"`
	Role          string `json:"role" binding:"required"`
}

// GetNonce 获取用于签名的nonce
func (h *AuthHandler) GetNonce(c *gin.Context) {
	var req GetNonceRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid wallet address",
		})
		return
	}

	// 验证钱包地址格式
	if !utils.IsValidEthereumAddress(req.WalletAddress) {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid Ethereum address format",
		})
		return
	}

	// 生成nonce
	nonce, err := utils.GenerateNonce()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to generate nonce",
		})
		return
	}

	// 查找或创建用户记录（仅用于存储nonce）
	var user models.User
	result := h.db.Where("wallet_address = ?", strings.ToLower(req.WalletAddress)).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		// 用户不存在，创建临时记录存储nonce
		user = models.User{
			WalletAddress: strings.ToLower(req.WalletAddress),
			Role:          "temp", // 临时角色
			Nonce:         nonce,
		}
		if err := h.db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Failed to create user record",
			})
			return
		}
	} else if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Database error",
		})
		return
	} else {
		// 更新现有用户的nonce
		user.Nonce = nonce
		if err := h.db.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Failed to update nonce",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"nonce": nonce,
		},
	})
}

// Login 用户登录
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request data",
		})
		return
	}

	// 验证钱包地址格式
	if !utils.IsValidEthereumAddress(req.WalletAddress) {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid Ethereum address format",
		})
		return
	}

	// 查找用户
	var user models.User
	result := h.db.Where("wallet_address = ?", strings.ToLower(req.WalletAddress)).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "User not found. Please register first.",
		})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Database error",
		})
		return
	}

	// 验证签名
	message := utils.GetSignMessage(user.Nonce)
	valid, err := utils.VerifySignature(req.WalletAddress, message, req.Signature)
	if err != nil || !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "Invalid signature",
		})
		return
	}

	// 如果用户角色是临时的，需要设置正确的角色
	if user.Role == "temp" {
		if req.Role == "" || !utils.IsValidRole(req.Role) {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Valid role required for first login",
			})
			return
		}
		user.Role = req.Role
		if err := h.db.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Failed to update user role",
			})
			return
		}
	}

	// 生成JWT token
	token, err := h.jwtManager.GenerateToken(user.ID, user.WalletAddress, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"token": token,
			"user":  user,
		},
	})
}

// Register 用户注册
func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request data",
		})
		return
	}

	// 验证输入
	if !utils.IsValidEthereumAddress(req.WalletAddress) {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid Ethereum address format",
		})
		return
	}

	if !utils.IsValidRole(req.Role) {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid role. Must be 'parent' or 'child'",
		})
		return
	}

	// 检查用户是否已存在
	var existingUser models.User
	result := h.db.Where("wallet_address = ?", strings.ToLower(req.WalletAddress)).First(&existingUser)
	if result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{
			"success": false,
			"error":   "User already exists",
		})
		return
	} else if result.Error != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Database error",
		})
		return
	}

	// 生成初始nonce
	nonce, err := utils.GenerateNonce()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to generate nonce",
		})
		return
	}

	// 创建新用户
	user := models.User{
		WalletAddress: strings.ToLower(req.WalletAddress),
		Role:          req.Role,
		Nonce:         nonce,
	}

	if err := h.db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data": gin.H{
			"user": user,
		},
	})
}

// Logout 用户登出
func (h *AuthHandler) Logout(c *gin.Context) {
	// 在实际应用中，这里可以将token加入黑名单
	// 目前只是返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Logged out successfully",
	})
}