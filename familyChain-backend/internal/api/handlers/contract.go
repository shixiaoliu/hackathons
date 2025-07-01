package handlers

import (
	"fmt"
	"math/big"
	"net/http"

	"eth-for-babies-backend/internal/services"
	"eth-for-babies-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ContractHandler struct {
	db              *gorm.DB
	contractService *services.ContractService
}

func NewContractHandler(db *gorm.DB, contractService *services.ContractService) *ContractHandler {
	return &ContractHandler{
		db:              db,
		contractService: contractService,
	}
}

type TransferRequest struct {
	To     string `json:"to" binding:"required"`
	Amount string `json:"amount" binding:"required"`
}

// GetBalance 获取代币余额
func (h *ContractHandler) GetBalance(c *gin.Context) {
	address := c.Param("address")

	// 验证地址格式
	if !utils.IsValidEthereumAddress(address) {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid Ethereum address format",
		})
		return
	}

	// 使用正确部署的RewardToken合约地址
	// 合约地址来自deployed_addresses.json
	tokenAddress := "0xe7cAa23a4E496e1A5854298Cd0b8f4Bd94C9F12d"

	// 从区块链获取真实代币余额
	balance, err := h.contractService.GetTokenBalance(tokenAddress, address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   fmt.Sprintf("Failed to get token balance: %v", err),
		})
		return
	}

	// 转换为浮点数字符串表示（假设代币有18位小数）
	divisor := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
	balanceFloat := new(big.Float).Quo(new(big.Float).SetInt(balance), divisor)
	balanceStr := balanceFloat.Text('f', 4) // 保留4位小数

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"balance": balanceStr,
		},
	})
}

// Transfer 转移代币
func (h *ContractHandler) Transfer(c *gin.Context) {
	var req TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request data",
		})
		return
	}

	// 验证地址格式
	if !utils.IsValidEthereumAddress(req.To) {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid recipient address format",
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

	// TODO: 实现实际的区块链转账
	// 这里返回模拟的交易哈希
	transactionHash := "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"transaction_hash": transactionHash,
			"from":             walletAddress,
			"to":               req.To,
			"amount":           req.Amount,
		},
	})
}

// GetTransactionStatus 获取交易状态
func (h *ContractHandler) GetTransactionStatus(c *gin.Context) {
	hash := c.Param("hash")

	// 简单验证哈希格式
	if len(hash) != 66 || hash[:2] != "0x" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid transaction hash format",
		})
		return
	}

	// TODO: 实现实际的交易状态查询
	// 这里返回模拟数据
	status := gin.H{
		"hash":          hash,
		"status":        "confirmed",
		"block_number":  12345678,
		"confirmations": 12,
		"gas_used":      "21000",
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    status,
	})
}
