package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

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

// NewExchangeHandler 创建新的兑换处理器
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

	// 创建一个带超时的上下文
	ctx, cancel := context.WithTimeout(c.Request.Context(), 60*time.Second) // 设置60秒超时
	defer cancel()

	// 记录开始处理时间
	startTime := time.Now()
	fmt.Printf("开始处理兑换请求，时间: %v, 奖品ID: %d, 用户: %s\n",
		startTime.Format(time.RFC3339), req.RewardID, child.Name)

	// 兑换奖品
	exchangeID, err := h.rewardService.ExchangeReward(ctx, child.ID, req)

	// 计算处理时间
	processingTime := time.Since(startTime)
	fmt.Printf("兑换请求处理完成，耗时: %v\n", processingTime)

	// 检查是否因为上下文超时而取消
	if ctx.Err() == context.DeadlineExceeded {
		fmt.Printf("兑换请求超时，处理时间超过60秒\n")
		c.JSON(http.StatusRequestTimeout, gin.H{
			"success": false,
			"error":   "请求超时，请稍后再试",
		})
		return
	}

	if err != nil {
		// 详细记录错误信息
		fmt.Printf("兑换奖品失败: %v, 奖品ID: %d, 用户: %s\n",
			err, req.RewardID, child.Name)

		// 根据错误类型返回不同的状态码
		statusCode := http.StatusInternalServerError
		errorMessage := "兑换奖品失败: " + err.Error()

		if strings.Contains(err.Error(), "insufficient funds") {
			statusCode = http.StatusBadRequest
			errorMessage = "余额不足，无法兑换奖品"
		} else if strings.Contains(err.Error(), "nonce too low") {
			errorMessage = "区块链交易错误，请稍后再试 (nonce错误)"
		}

		c.JSON(statusCode, gin.H{
			"success": false,
			"error":   errorMessage,
		})
		return
	}

	// 获取兑换详情
	exchange, err := h.rewardService.GetExchange(ctx, exchangeID)
	if err != nil {
		fmt.Printf("获取兑换详情失败: %v, 兑换ID: %d\n", err, exchangeID)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "获取兑换详情失败: " + err.Error(),
		})
		return
	}

	fmt.Printf("兑换成功，兑换ID: %d, 奖品: %s, 用户: %s\n",
		exchangeID, exchange.RewardName, child.Name)

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
		fmt.Printf("获取兑换记录失败: 用户未认证\n")
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "用户未认证",
		})
		return
	}

	fmt.Printf("获取孩子兑换记录，childID: %v, 类型: %T\n", childID, childID)

	// 获取兑换记录
	exchanges, err := h.rewardService.GetChildExchanges(c.Request.Context(), childID.(uint))
	if err != nil {
		fmt.Printf("获取兑换记录失败: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "获取兑换记录失败: " + err.Error(),
		})
		return
	}

	fmt.Printf("成功获取到 %d 条兑换记录\n", len(exchanges))
	for i, exchange := range exchanges {
		fmt.Printf("兑换记录 %d: ID=%d, 奖品=%s, 状态=%s\n", i+1, exchange.ID, exchange.RewardName, exchange.Status)
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
