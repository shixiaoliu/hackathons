package handlers

import (
	"fmt"
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
		fmt.Printf("请求解析失败: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "无效的请求数据: " + err.Error(),
		})
		return
	}

	// 记录请求数据
	fmt.Printf("收到创建奖品请求: %+v\n", req)

	// 验证请求数据
	if req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "奖品名称不能为空",
		})
		return
	}

	// 验证图片URL
	if req.ImageURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "奖品图片不能为空",
		})
		return
	}

	// 确保价格有效
	if req.TokenPrice <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "奖品价格必须大于0",
		})
		return
	}

	// 确保库存有效
	if req.Stock < 0 {
		req.Stock = 0
	}

	// 获取当前用户信息
	userID, exists := c.Get("user_id")
	if !exists {
		fmt.Println("用户未认证或获取用户ID失败")
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "用户未认证",
		})
		return
	}
	fmt.Printf("当前用户ID: %v, 类型: %T\n", userID, userID)

	// 获取家庭ID
	familyIDStr := c.Param("family_id")
	fmt.Printf("请求中的家庭ID参数: %s\n", familyIDStr)

	familyID, err := strconv.ParseUint(familyIDStr, 10, 64)
	if err != nil {
		fmt.Printf("家庭ID解析失败: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "无效的家庭ID",
		})
		return
	}

	// 净化输入数据
	req.Name = utils.SanitizeString(req.Name)
	req.Description = utils.SanitizeString(req.Description)
	fmt.Printf("净化后的数据: 名称=%s, 描述=%s\n", req.Name, req.Description)

	// 创建奖品
	fmt.Printf("调用service层创建奖品, 用户ID=%d, 家庭ID=%d\n", userID.(uint), uint(familyID))
	rewardID, err := h.rewardService.CreateReward(c.Request.Context(), userID.(uint), uint(familyID), req)
	if err != nil {
		fmt.Printf("创建奖品失败: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "创建奖品失败: " + err.Error(),
		})
		return
	}
	fmt.Printf("奖品创建成功，ID: %d\n", rewardID)

	// 获取创建的奖品详情
	reward, err := h.rewardService.GetReward(c.Request.Context(), rewardID)
	if err != nil {
		fmt.Printf("获取创建的奖品详情失败: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "获取创建的奖品失败",
		})
		return
	}
	fmt.Printf("返回奖品详情: %+v\n", reward)

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
		fmt.Printf("更新奖品请求解析失败: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "无效的请求数据: " + err.Error(),
		})
		return
	}

	fmt.Printf("收到更新奖品请求: %+v\n", req)

	// 获取奖品ID
	rewardIDStr := c.Param("id")
	fmt.Printf("奖品ID参数: %s\n", rewardIDStr)

	rewardID, err := strconv.ParseUint(rewardIDStr, 10, 64)
	if err != nil {
		fmt.Printf("奖品ID解析失败: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "无效的奖品ID",
		})
		return
	}

	// 获取当前用户信息
	userID, exists := c.Get("user_id")
	if !exists {
		fmt.Println("用户未认证或获取用户ID失败")
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "用户未认证",
		})
		return
	}
	fmt.Printf("当前用户ID: %v\n", userID)

	// 更新奖品
	fmt.Printf("准备更新奖品: ID=%d\n", uint(rewardID))
	if req.Name != nil && *req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "奖品名称不能为空",
		})
		return
	}

	if req.ImageURL != nil && *req.ImageURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "奖品图片不能为空",
		})
		return
	}

	if req.TokenPrice != nil && *req.TokenPrice <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "奖品价格必须大于0",
		})
		return
	}

	// 净化输入数据
	if req.Name != nil {
		sanitizedName := utils.SanitizeString(*req.Name)
		req.Name = &sanitizedName
	}
	if req.Description != nil {
		sanitizedDesc := utils.SanitizeString(*req.Description)
		req.Description = &sanitizedDesc
	}
	fmt.Printf("处理后的请求数据: %+v\n", req)

	// 更新奖品
	fmt.Printf("调用service层更新奖品, ID=%d\n", rewardID)
	err = h.rewardService.UpdateReward(c.Request.Context(), uint(rewardID), req)
	if err != nil {
		fmt.Printf("更新奖品失败: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "更新奖品失败: " + err.Error(),
		})
		return
	}

	// 获取更新后的奖品详情
	reward, err := h.rewardService.GetReward(c.Request.Context(), uint(rewardID))
	if err != nil {
		fmt.Printf("获取更新后的奖品失败: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "获取更新后的奖品失败",
		})
		return
	}
	fmt.Printf("奖品更新成功: %+v\n", reward)

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
