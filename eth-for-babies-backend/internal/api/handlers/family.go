package handlers

import (
	"net/http"
	"strconv"

	"eth-for-babies-backend/internal/models"
	"eth-for-babies-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FamilyHandler struct {
	db *gorm.DB
}

func NewFamilyHandler(db *gorm.DB) *FamilyHandler {
	return &FamilyHandler{db: db}
}

type CreateFamilyRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateFamilyRequest struct {
	Name string `json:"name"`
}

type AddMemberRequest struct {
	WalletAddress string `json:"wallet_address" binding:"required"`
}

// CreateFamily 创建家庭
func (h *FamilyHandler) CreateFamily(c *gin.Context) {
	var req CreateFamilyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request data",
		})
		return
	}

	// 获取当前用户信息
	walletAddress, exists := c.Get("wallet_address")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "User not authenticated",
		})
		return
	}

	role, exists := c.Get("role")
	if !exists || role != "parent" {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "Only parents can create families",
		})
		return
	}

	// 检查是否已经有家庭
	var existingFamily models.Family
	result := h.db.Where("parent_address = ?", walletAddress).First(&existingFamily)
	if result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{
			"success": false,
			"error":   "Family already exists for this parent",
		})
		return
	} else if result.Error != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Database error",
		})
		return
	}

	// 创建家庭
	family := models.Family{
		Name:          utils.SanitizeString(req.Name),
		ParentAddress: walletAddress.(string),
	}

	if err := h.db.Create(&family).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to create family",
		})
		return
	}

	// 预加载关联数据
	h.db.Preload("Children").First(&family, family.ID)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    family,
	})
}

// GetFamilies 获取家庭列表
func (h *FamilyHandler) GetFamilies(c *gin.Context) {
	walletAddress, exists := c.Get("wallet_address")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "User not authenticated",
		})
		return
	}

	role, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "User role not found",
		})
		return
	}

	var families []models.Family
	var query *gorm.DB

	if role == "parent" {
		// 父母只能看到自己的家庭
		query = h.db.Where("parent_address = ?", walletAddress)
	} else {
		// 孩子可以看到自己所属的家庭
		query = h.db.Joins("JOIN children ON families.parent_address = children.parent_address").Where("children.wallet_address = ?", walletAddress)
	}

	if err := query.Preload("Children").Find(&families).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to fetch families",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    families,
	})
}

// GetFamilyByID 获取家庭详情
func (h *FamilyHandler) GetFamilyByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid family ID",
		})
		return
	}

	walletAddress, exists := c.Get("wallet_address")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "User not authenticated",
		})
		return
	}

	var family models.Family
	result := h.db.Preload("Children").First(&family, uint(id))
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Family not found",
		})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Database error",
		})
		return
	}

	// 检查权限
	role, _ := c.Get("role")
	if role == "parent" && family.ParentAddress != walletAddress.(string) {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "Access denied",
		})
		return
	} else if role == "child" {
		// 检查孩子是否属于这个家庭
		var child models.Child
		result := h.db.Where("wallet_address = ? AND parent_address = ?", walletAddress, family.ParentAddress).First(&child)
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"error":   "Access denied",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    family,
	})
}

// UpdateFamily 更新家庭信息
func (h *FamilyHandler) UpdateFamily(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid family ID",
		})
		return
	}

	var req UpdateFamilyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request data",
		})
		return
	}

	walletAddress, exists := c.Get("wallet_address")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "User not authenticated",
		})
		return
	}

	// 查找家庭
	var family models.Family
	result := h.db.First(&family, uint(id))
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Family not found",
		})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Database error",
		})
		return
	}

	// 检查权限（只有家庭的父母可以更新）
	if family.ParentAddress != walletAddress.(string) {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "Only the family parent can update family information",
		})
		return
	}

	// 更新家庭信息
	if req.Name != "" {
		family.Name = utils.SanitizeString(req.Name)
	}

	if err := h.db.Save(&family).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to update family",
		})
		return
	}

	// 预加载关联数据
	h.db.Preload("Children").First(&family, family.ID)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    family,
	})
}