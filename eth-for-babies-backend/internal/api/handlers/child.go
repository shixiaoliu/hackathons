package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"eth-for-babies-backend/internal/models"
	"eth-for-babies-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ChildHandler struct {
	db *gorm.DB
}

func NewChildHandler(db *gorm.DB) *ChildHandler {
	return &ChildHandler{db: db}
}

type CreateChildRequest struct {
	Name          string `json:"name" binding:"required"`
	WalletAddress string `json:"wallet_address" binding:"required"`
	Age           int    `json:"age" binding:"required,min=1,max=18"`
	Avatar        string `json:"avatar,omitempty"`
}

type UpdateChildRequest struct {
	Name   string `json:"name,omitempty"`
	Age    int    `json:"age,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}

// CreateChild 添加孩子
func (h *ChildHandler) CreateChild(c *gin.Context) {
	var req CreateChildRequest
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
			"error":   "Only parents can add children",
		})
		return
	}

	// 验证孩子的钱包地址格式
	if !utils.IsValidEthereumAddress(req.WalletAddress) {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid Ethereum address format",
		})
		return
	}

	// 检查孩子是否已存在
	var existingChild models.Child
	result := h.db.Where("wallet_address = ?", strings.ToLower(req.WalletAddress)).First(&existingChild)
	if result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{
			"success": false,
			"error":   "Child with this wallet address already exists",
		})
		return
	} else if result.Error != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Database error",
		})
		return
	}

	// 确保父母有家庭
	var family models.Family
	result = h.db.Where("parent_address = ?", walletAddress).First(&family)
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Please create a family first",
		})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Database error",
		})
		return
	}

	// 创建孩子记录
	child := models.Child{
		Name:                utils.SanitizeString(req.Name),
		WalletAddress:       strings.ToLower(req.WalletAddress),
		Age:                 req.Age,
		ParentAddress:       walletAddress.(string),
		TotalTasksCompleted: 0,
		TotalRewardsEarned:  "0",
	}

	if req.Avatar != "" {
		child.Avatar = &req.Avatar
	}

	if err := h.db.Create(&child).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to create child",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    child,
	})
}

// GetChildren 获取孩子列表
func (h *ChildHandler) GetChildren(c *gin.Context) {
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

	var children []models.Child
	var query *gorm.DB

	if role == "parent" {
		// 父母可以看到自己的所有孩子
		query = h.db.Where("parent_address = ?", walletAddress)
	} else {
		// 孩子只能看到自己的信息
		query = h.db.Where("wallet_address = ?", walletAddress)
	}

	if err := query.Find(&children).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to fetch children",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    children,
	})
}

// GetChildByID 获取孩子详情
func (h *ChildHandler) GetChildByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid child ID",
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

	var child models.Child
	result := h.db.First(&child, uint(id))
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Child not found",
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
	if role == "parent" && child.ParentAddress != walletAddress.(string) {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "Access denied",
		})
		return
	} else if role == "child" && child.WalletAddress != walletAddress.(string) {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "Access denied",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    child,
	})
}

// UpdateChild 更新孩子信息
func (h *ChildHandler) UpdateChild(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid child ID",
		})
		return
	}

	var req UpdateChildRequest
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

	// 查找孩子
	var child models.Child
	result := h.db.First(&child, uint(id))
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Child not found",
		})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Database error",
		})
		return
	}

	// 检查权限（父母或孩子本人可以更新）
	role, _ := c.Get("role")
	if role == "parent" && child.ParentAddress != walletAddress.(string) {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "Access denied",
		})
		return
	} else if role == "child" && child.WalletAddress != walletAddress.(string) {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "Access denied",
		})
		return
	}

	// 更新孩子信息
	if req.Name != "" {
		child.Name = utils.SanitizeString(req.Name)
	}
	if req.Age > 0 {
		child.Age = req.Age
	}
	if req.Avatar != "" {
		child.Avatar = &req.Avatar
	}

	if err := h.db.Save(&child).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to update child",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    child,
	})
}

// GetChildProgress 获取孩子进度
func (h *ChildHandler) GetChildProgress(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid child ID",
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

	// 查找孩子
	var child models.Child
	result := h.db.First(&child, uint(id))
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Child not found",
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
	if role == "parent" && child.ParentAddress != walletAddress.(string) {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "Access denied",
		})
		return
	} else if role == "child" && child.WalletAddress != walletAddress.(string) {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "Access denied",
		})
		return
	}

	// 获取任务统计
	var taskStats struct {
		Pending    int64 `json:"pending"`
		InProgress int64 `json:"in_progress"`
		Completed  int64 `json:"completed"`
		Approved   int64 `json:"approved"`
		Rejected   int64 `json:"rejected"`
	}

	h.db.Model(&models.Task{}).Where("assigned_child_id = ? AND status = ?", id, "pending").Count(&taskStats.Pending)
	h.db.Model(&models.Task{}).Where("assigned_child_id = ? AND status = ?", id, "in_progress").Count(&taskStats.InProgress)
	h.db.Model(&models.Task{}).Where("assigned_child_id = ? AND status = ?", id, "completed").Count(&taskStats.Completed)
	h.db.Model(&models.Task{}).Where("assigned_child_id = ? AND status = ?", id, "approved").Count(&taskStats.Approved)
	h.db.Model(&models.Task{}).Where("assigned_child_id = ? AND status = ?", id, "rejected").Count(&taskStats.Rejected)

	progress := gin.H{
		"child":                 child,
		"total_tasks_completed": child.TotalTasksCompleted,
		"total_rewards_earned":  child.TotalRewardsEarned,
		"task_statistics":       taskStats,
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    progress,
	})
}

// DeleteChild 删除孩子
func (h *ChildHandler) DeleteChild(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid child ID",
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

	// 查找孩子
	var child models.Child
	result := h.db.First(&child, uint(id))
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Child not found",
		})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Database error",
		})
		return
	}

	// 只有父母能删除自己的孩子
	role, _ := c.Get("role")
	if role != "parent" || child.ParentAddress != walletAddress.(string) {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "Access denied",
		})
		return
	}

	if err := h.db.Delete(&child).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to delete child",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Child deleted successfully",
	})
}
