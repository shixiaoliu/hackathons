package handlers

import (
	"net/http"
	"strconv"

	"eth-for-babies-backend/internal/models"
	"eth-for-babies-backend/internal/services"
	"eth-for-babies-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

// RewardHandler 处理奖品相关的API请求
type RewardHandler struct {
	rewardService *services.RewardService
}

// NewRewardHandler 创建一个新的奖品处理器
func NewRewardHandler(rewardService *services.RewardService) *RewardHandler {
	return &RewardHandler{
		rewardService: rewardService,
	}
}

// CreateReward 创建新的实物奖励
func (h *RewardHandler) CreateReward(c *gin.Context) {
	var req models.RewardCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "无效的请求数据: " + err.Error(),
		})
		return
	}

	// 获取当前用户信息
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "用户未认证",
		})
		return
	}

	// 获取家庭ID
	familyIDStr := c.Param("family_id")
	familyID, err := strconv.ParseUint(familyIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "无效的家庭ID",
		})
		return
	}

	// 净化输入数据
	req.Name = utils.SanitizeString(req.Name)
	req.Description = utils.SanitizeString(req.Description)

	// 创建奖品
	rewardID, err := h.rewardService.CreateReward(c.Request.Context(), userID.(uint), uint(familyID), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "创建奖品失败: " + err.Error(),
		})
		return
	}

	// 获取创建的奖品详情
	reward, err := h.rewardService.GetReward(c.Request.Context(), rewardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "获取创建的奖品失败",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    reward,
	})
}

// GetRewards 获取家庭奖品列表
func (h *RewardHandler) GetRewards(c *gin.Context) {
	// 获取家庭ID
	familyIDStr := c.Param("family_id")
	familyID, err := strconv.ParseUint(familyIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "无效的家庭ID",
		})
		return
	}

	// 获取查询参数
	activeOnly := true
	if c.Query("active_only") == "false" {
		activeOnly = false
	}

	// 获取奖品列表
	rewards, err := h.rewardService.GetFamilyRewards(c.Request.Context(), uint(familyID), activeOnly)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "获取奖品列表失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    rewards,
	})
}

// GetRewardByID 获取奖品详情
func (h *RewardHandler) GetRewardByID(c *gin.Context) {
	// 获取奖品ID
	rewardIDStr := c.Param("id")
	rewardID, err := strconv.ParseUint(rewardIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "无效的奖品ID",
		})
		return
	}

	// 获取奖品详情
	reward, err := h.rewardService.GetReward(c.Request.Context(), uint(rewardID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "获取奖品详情失败: " + err.Error(),
		})
		return
	}

	if reward == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "奖品不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    reward,
	})
}

// UpdateReward 更新奖品信息
func (h *RewardHandler) UpdateReward(c *gin.Context) {
	var req models.RewardUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "无效的请求数据: " + err.Error(),
		})
		return
	}

	// 获取奖品ID
	rewardIDStr := c.Param("id")
	rewardID, err := strconv.ParseUint(rewardIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "无效的奖品ID",
		})
		return
	}

	// 获取当前用户信息
	_, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "用户未认证",
		})
		return
	}

	// 净化输入数据
	if req.Name != "" {
		req.Name = utils.SanitizeString(req.Name)
	}
	if req.Description != "" {
		req.Description = utils.SanitizeString(req.Description)
	}

	// 更新奖品
	err = h.rewardService.UpdateReward(c.Request.Context(), uint(rewardID), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "更新奖品失败: " + err.Error(),
		})
		return
	}

	// 获取更新后的奖品详情
	reward, err := h.rewardService.GetReward(c.Request.Context(), uint(rewardID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "获取更新后的奖品失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    reward,
	})
}

// DeleteReward 删除奖品
func (h *RewardHandler) DeleteReward(c *gin.Context) {
	// 获取奖品ID
	rewardIDStr := c.Param("id")
	rewardID, err := strconv.ParseUint(rewardIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "无效的奖品ID",
		})
		return
	}

	// 获取当前用户信息
	_, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "用户未认证",
		})
		return
	}

	// 更新奖品状态为非活跃
	active := false
	req := models.RewardUpdateRequest{
		Active: &active,
	}

	err = h.rewardService.UpdateReward(c.Request.Context(), uint(rewardID), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "删除奖品失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "奖品已成功删除",
	})
}
