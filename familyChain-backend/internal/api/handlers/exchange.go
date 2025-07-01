package handlers

import (
	"net/http"
	"strconv"

	"eth-for-babies-backend/internal/models"
	"eth-for-babies-backend/internal/services"
	"eth-for-babies-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

// ExchangeHandler 处理兑换相关的API请求
type ExchangeHandler struct {
	rewardService *services.RewardService
	childService  *services.ChildService
}

// NewExchangeHandler 创建一个新的兑换处理器
func NewExchangeHandler(rewardService *services.RewardService, childService *services.ChildService) *ExchangeHandler {
	return &ExchangeHandler{
		rewardService: rewardService,
		childService:  childService,
	}
}

// ExchangeReward 兑换奖品
func (h *ExchangeHandler) ExchangeReward(c *gin.Context) {
	var req models.ExchangeCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "无效的请求数据: " + err.Error(),
		})
		return
	}

	// 检查用户认证
	if _, exists := c.Get("user_id"); !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "用户未认证",
		})
		return
	}

	// 获取用户钱包地址
	walletAddress, exists := c.Get("wallet_address")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "用户钱包地址未找到",
		})
		return
	}

	// 通过钱包地址获取child信息
	child, err := h.childService.GetByWalletAddress(c.Request.Context(), walletAddress.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "获取孩子信息失败: " + err.Error(),
		})
		return
	}
	if child == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "未找到对应的孩子记录",
		})
		return
	}

	// 净化输入数据
	req.Notes = utils.SanitizeString(req.Notes)

	// 兑换奖品
	exchangeID, err := h.rewardService.ExchangeReward(c.Request.Context(), child.ID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "兑换奖品失败: " + err.Error(),
		})
		return
	}

	// 获取兑换详情
	exchange, err := h.rewardService.GetExchange(c.Request.Context(), exchangeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "获取兑换详情失败",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    exchange,
	})
}

// GetChildExchanges 获取孩子的兑换记录
func (h *ExchangeHandler) GetChildExchanges(c *gin.Context) {
	// 获取当前用户信息
	childID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "用户未认证",
		})
		return
	}

	// 获取兑换记录
	exchanges, err := h.rewardService.GetChildExchanges(c.Request.Context(), childID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "获取兑换记录失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    exchanges,
	})
}

// GetExchangeByID 获取兑换详情
func (h *ExchangeHandler) GetExchangeByID(c *gin.Context) {
	// 获取兑换ID
	exchangeIDStr := c.Param("id")
	exchangeID, err := strconv.ParseUint(exchangeIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "无效的兑换ID",
		})
		return
	}

	// 获取兑换详情
	exchange, err := h.rewardService.GetExchange(c.Request.Context(), uint(exchangeID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "获取兑换详情失败: " + err.Error(),
		})
		return
	}

	if exchange == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "兑换记录不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    exchange,
	})
}

// GetFamilyExchanges 获取家庭的兑换记录
func (h *ExchangeHandler) GetFamilyExchanges(c *gin.Context) {
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

	// 获取兑换记录
	exchanges, err := h.rewardService.GetFamilyExchanges(c.Request.Context(), uint(familyID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "获取兑换记录失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    exchanges,
	})
}

// UpdateExchangeStatus 更新兑换状态
func (h *ExchangeHandler) UpdateExchangeStatus(c *gin.Context) {
	var req models.ExchangeUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "无效的请求数据: " + err.Error(),
		})
		return
	}

	// 获取兑换ID
	exchangeIDStr := c.Param("id")
	exchangeID, err := strconv.ParseUint(exchangeIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "无效的兑换ID",
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
	req.Notes = utils.SanitizeString(req.Notes)

	// 更新兑换状态
	err = h.rewardService.UpdateExchangeStatus(c.Request.Context(), uint(exchangeID), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "更新兑换状态失败: " + err.Error(),
		})
		return
	}

	// 获取更新后的兑换详情
	exchange, err := h.rewardService.GetExchange(c.Request.Context(), uint(exchangeID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "获取更新后的兑换详情失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    exchange,
	})
}
