package handlers

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"eth-for-babies-backend/internal/models"
	"eth-for-babies-backend/internal/utils"
	"eth-for-babies-backend/pkg/blockchain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TaskHandler struct {
	db              *gorm.DB
	contractManager *blockchain.ContractManager
}

func NewTaskHandler(db *gorm.DB, contractManager *blockchain.ContractManager) *TaskHandler {
	return &TaskHandler{
		db:              db,
		contractManager: contractManager,
	}
}

type CreateTaskRequest struct {
	Title           string  `json:"title" binding:"required"`
	Description     string  `json:"description" binding:"required"`
	RewardAmount    string  `json:"reward_amount" binding:"required"`
	Difficulty      string  `json:"difficulty" binding:"required"`
	AssignedChildID *uint   `json:"assigned_child_id,omitempty"`
	DueDate         string  `json:"due_date,omitempty"`
	ContractTaskID  *uint64 `json:"contract_task_id,omitempty"`
}

type UpdateTaskRequest struct {
	Title           string `json:"title,omitempty"`
	Description     string `json:"description,omitempty"`
	RewardAmount    string `json:"reward_amount,omitempty"`
	Difficulty      string `json:"difficulty,omitempty"`
	Status          string `json:"status,omitempty"`
	AssignedChildID *uint  `json:"assigned_child_id,omitempty"`
	DueDate         string `json:"due_date,omitempty"`
}

type CompleteTaskRequest struct {
	CompletionProof string `json:"completion_proof" binding:"required"`
}

type RejectTaskRequest struct {
	Reason string `json:"reason,omitempty"`
}

// CreateTask 创建任务
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req CreateTaskRequest
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
			"error":   "Only parents can create tasks",
		})
		return
	}

	// 验证难度
	if !utils.IsValidDifficulty(req.Difficulty) {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid difficulty. Must be 'easy', 'medium', or 'hard'",
		})
		return
	}

	// 如果指定了孩子，验证孩子是否属于当前父母
	if req.AssignedChildID != nil {
		var child models.Child
		result := h.db.Where("id = ? AND parent_address = ?", *req.AssignedChildID, walletAddress).First(&child)
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Child not found or not belongs to you",
			})
			return
		} else if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Database error",
			})
			return
		}
	}

	// 创建任务
	task := models.Task{
		Title:           utils.SanitizeString(req.Title),
		Description:     utils.SanitizeString(req.Description),
		RewardAmount:    req.RewardAmount,
		Difficulty:      req.Difficulty,
		Status:          "pending",
		CreatedBy:       walletAddress.(string),
		AssignedChildID: req.AssignedChildID,
	}

	// 设置合约任务ID（如果前端提供了）
	if req.ContractTaskID != nil {
		task.ContractTaskID = req.ContractTaskID
	}

	// 解析截止日期
	if req.DueDate != "" {
		dueDate, err := time.Parse(time.RFC3339, req.DueDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Invalid due date format. Use RFC3339 format",
			})
			return
		}
		task.DueDate = &dueDate
	}

	// 如果分配给了孩子，状态改为进行中
	if req.AssignedChildID != nil {
		log.Printf("assigned child id: %d", *req.AssignedChildID)
		task.Status = "in_progress"
	}

	// 保存到数据库
	if err := h.db.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to create task",
		})
		return
	}

	// 预加载关联数据
	h.db.Preload("AssignedChild").First(&task, task.ID)

	// 返回响应
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    task,
	})
}

// GetTasks 获取任务列表
func (h *TaskHandler) GetTasks(c *gin.Context) {
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

	// 获取查询参数
	childIDParam := c.Query("child_id")
	status := c.Query("status")

	var tasks []models.Task
	query := h.db.Model(&models.Task{})

	if role == "parent" {
		// 父母可以看到自己创建的所有任务
		query = query.Where("created_by = ?", walletAddress)

		// 如果指定了孩子ID，过滤该孩子的任务
		if childIDParam != "" {
			childID, err := strconv.ParseUint(childIDParam, 10, 32)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"success": false,
					"error":   "Invalid child ID",
				})
				return
			}
			query = query.Where("assigned_child_id = ?", uint(childID))
		}
	} else {
		// 孩子只能看到分配给自己的任务
		var child models.Child
		result := h.db.Where("wallet_address = ?", walletAddress).First(&child)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Child record not found",
			})
			return
		}
		query = query.Where("assigned_child_id = ?", child.ID)
	}

	// 按状态过滤
	if status != "" && utils.IsValidTaskStatus(status) {
		query = query.Where("status = ?", status)
	}

	// 执行查询
	if err := query.Preload("AssignedChild").Order("created_at DESC").Find(&tasks).Error; err != nil {
		log.Println("查询任务时出错:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to fetch tasks",
		})
		return
	}

	log.Printf("获取到的任务数量: %d", len(tasks)) // 添加调试信息
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    tasks,
	})
}

// GetTaskByID 获取任务详情
func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid task ID",
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

	var task models.Task
	result := h.db.Preload("AssignedChild").First(&task, uint(id))
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Task not found",
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
	if role == "parent" && task.CreatedBy != walletAddress.(string) {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "Access denied",
		})
		return
	} else if role == "child" {
		// 检查任务是否分配给了这个孩子
		var child models.Child
		result := h.db.Where("wallet_address = ?", walletAddress).First(&child)
		if result.Error != nil || task.AssignedChildID == nil || *task.AssignedChildID != child.ID {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"error":   "Access denied",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    task,
	})
}

// UpdateTask 更新任务
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid task ID",
		})
		return
	}

	var req UpdateTaskRequest
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

	// 查找任务
	var task models.Task
	result := h.db.First(&task, uint(id))
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Task not found",
		})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Database error",
		})
		return
	}

	// 检查权限（只有任务创建者可以更新）
	if task.CreatedBy != walletAddress.(string) {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "Only task creator can update the task",
		})
		return
	}

	// 更新任务信息
	if req.Title != "" {
		task.Title = utils.SanitizeString(req.Title)
	}
	if req.Description != "" {
		task.Description = utils.SanitizeString(req.Description)
	}
	if req.RewardAmount != "" {
		task.RewardAmount = req.RewardAmount
	}
	if req.Difficulty != "" && utils.IsValidDifficulty(req.Difficulty) {
		task.Difficulty = req.Difficulty
	}
	if req.Status != "" && utils.IsValidTaskStatus(req.Status) {
		task.Status = req.Status
	}
	if req.AssignedChildID != nil {
		// 验证孩子是否属于当前父母
		var child models.Child
		result := h.db.Where("id = ? AND parent_address = ?", *req.AssignedChildID, walletAddress).First(&child)
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Child not found or not belongs to you",
			})
			return
		} else if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Database error",
			})
			return
		}

		// 更新任务分配
		task.AssignedChildID = req.AssignedChildID

		// 如果分配给了孩子且状态是pending，改为in_progress
		if task.Status == "pending" {
			task.Status = "in_progress"
		}
	}
	if req.DueDate != "" {
		dueDate, err := time.Parse(time.RFC3339, req.DueDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Invalid due date format. Use RFC3339 format",
			})
			return
		}
		task.DueDate = &dueDate
	}

	if err := h.db.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to update task",
		})
		return
	}

	// 预加载关联数据
	h.db.Preload("AssignedChild").First(&task, task.ID)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    task,
	})
}

// CompleteTask 完成任务
func (h *TaskHandler) CompleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid task ID",
		})
		return
	}

	var req CompleteTaskRequest
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

	role, exists := c.Get("role")
	if !exists || role != "child" {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "Only children can complete tasks",
		})
		return
	}

	// 查找任务
	var task models.Task
	result := h.db.First(&task, uint(id))
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Task not found",
		})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Database error",
		})
		return
	}

	// 检查任务是否分配给了当前孩子
	var child models.Child
	result = h.db.Where("wallet_address = ?", walletAddress).First(&child)
	if result.Error != nil || task.AssignedChildID == nil || *task.AssignedChildID != child.ID {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "Task not assigned to you",
		})
		return
	}

	// 检查任务状态
	if task.Status != "in_progress" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Task is not in progress",
		})
		return
	}

	// 更新任务状态
	now := time.Now()
	task.Status = "completed"
	task.CompletionProof = &req.CompletionProof
	task.SubmittedAt = &now

	if err := h.db.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to complete task",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    task,
	})
}

// ApproveTask 批准任务
func (h *TaskHandler) ApproveTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid task ID",
		})
		return
	}

	// 输出合约管理器和合约配置信息
	log.Printf("[DEBUG] ===== 合约配置诊断 =====")
	if h.contractManager == nil {
		log.Printf("[ERROR] 合约管理器未初始化")
	} else {
		addresses := h.contractManager.GetContractAddresses()
		log.Printf("[DEBUG] 合约地址配置:")
		log.Printf("[DEBUG]   - 任务合约: %s", addresses["task"])
		log.Printf("[DEBUG]   - 家庭合约: %s", addresses["family"])
		log.Printf("[DEBUG]   - 代币合约: %s", addresses["token"])
		log.Printf("[DEBUG]   - 奖励合约: %s", addresses["reward"])

		// 检查是否存在零地址
		zeroAddress := "0x0000000000000000000000000000000000000000"
		if addresses["token"] == zeroAddress || addresses["token"] == "" {
			log.Printf("[ERROR] 代币合约地址为空或零地址，请检查环境变量 TOKEN_CONTRACT_ADDRESS")
		}
	}

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
			"error":   "Only parents can approve tasks",
		})
		return
	}

	// 查找任务
	var task models.Task
	result := h.db.Preload("AssignedChild").First(&task, uint(id))
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Task not found",
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
	if task.CreatedBy != walletAddress.(string) {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "Only task creator can approve the task",
		})
		return
	}

	// 检查任务状态
	if task.Status != "completed" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Task is not completed yet",
		})
		return
	}

	// 开始事务
	tx := h.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 更新任务状态
	now := time.Now()
	task.Status = "approved"
	task.ApprovedAt = &now

	if err := tx.Save(&task).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to approve task",
		})
		return
	}

	// 更新孩子的统计信息
	if task.AssignedChild != nil {
		if err := tx.Model(&models.Child{}).Where("id = ?", task.AssignedChild.ID).Updates(map[string]interface{}{
			"total_tasks_completed": gorm.Expr("total_tasks_completed + ?", 1),
			"total_rewards_earned":  gorm.Expr("CAST(total_rewards_earned AS REAL) + CAST(? AS REAL)", task.RewardAmount),
		}).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Failed to update child statistics",
			})
			return
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to commit transaction",
		})
		return
	}

	// 铸造代币奖励给孩子（数量是ETH奖励的10000倍）
	log.Printf("[DEBUG] 尝试铸造代币奖励，任务ID: %d", id)

	if task.AssignedChild == nil {
		log.Printf("[DEBUG] 没有分配的孩子，无法铸造代币")
	} else if task.AssignedChild.WalletAddress == "" {
		log.Printf("[DEBUG] 孩子(%d)没有钱包地址，无法铸造代币", task.AssignedChild.ID)
	} else if h.contractManager == nil {
		log.Printf("[DEBUG] 合约管理器未初始化，无法铸造代币")
	} else if h.contractManager.RewardToken == nil {
		log.Printf("[DEBUG] 奖励代币合约未初始化，尝试手动初始化...")

		// 获取代币合约地址
		tokenAddress := h.contractManager.GetContractAddresses()["token"]
		if tokenAddress == "" || tokenAddress == "0x0000000000000000000000000000000000000000" {
			log.Printf("[ERROR] 无法初始化代币合约：地址为空或零地址")
		} else {
			log.Printf("[DEBUG] 尝试使用地址初始化代币合约: %s", tokenAddress)

			// 尝试手动初始化RewardToken合约
			err := h.initializeRewardToken(tokenAddress)
			if err != nil {
				log.Printf("[ERROR] 手动初始化代币合约失败: %v", err)
			} else {
				log.Printf("[SUCCESS] 手动初始化代币合约成功！")
			}
		}
	} else {
		log.Printf("[DEBUG] 所有条件满足，开始铸造代币流程")

		// 解析奖励金额
		log.Printf("[DEBUG] 原始奖励金额: %s", task.RewardAmount)
		rewardFloat, err := strconv.ParseFloat(task.RewardAmount, 64)
		if err != nil {
			log.Printf("[ERROR] 无法解析奖励金额：%v", err)
		} else {
			// 将奖励乘以10000并考虑代币小数位
			tokenAmountFloat := rewardFloat * 10000 * 1e18 // 乘以10^18考虑代币的18个小数位

			// 使用math/big处理大数，避免整数溢出
			tokenAmountBig := new(big.Float).SetFloat64(tokenAmountFloat)
			tokenAmountInt, _ := tokenAmountBig.Int(nil)

			log.Printf("[DEBUG] 计算的代币金额: %s (原始奖励: %f × 10000 × 10^18)", tokenAmountInt.String(), rewardFloat)
			log.Printf("[DEBUG] 这相当于 %f 个完整代币单位", rewardFloat*10000)

			// 使用GetTransactOpts替代client.GetAuth()
			log.Printf("[DEBUG] 尝试获取交易选项...")
			ctx := context.Background()
			auth, err := h.contractManager.GetTransactOpts(ctx)
			if err != nil {
				log.Printf("[ERROR] 获取交易选项失败：%v", err)
			} else {
				log.Printf("[DEBUG] 成功获取交易选项，使用地址: %s", auth.From.Hex())

				// 执行铸币操作
				childAddress := common.HexToAddress(task.AssignedChild.WalletAddress)
				log.Printf("[DEBUG] 孩子钱包地址: %s", childAddress.Hex())

				log.Printf("[DEBUG] 开始调用RewardToken.Mint方法...")

				// 检查RewardToken合约初始化状态
				log.Printf("[DEBUG] RewardToken合约地址: %s", h.contractManager.GetContractAddresses()["token"])

				// 尝试调用简单的查询方法来验证合约连接
				symbol, err := h.contractManager.RewardToken.Symbol(nil)
				if err != nil {
					log.Printf("[ERROR] RewardToken合约连接测试失败: %v", err)
					log.Printf("[DEBUG] 尝试继续铸币流程...")
				} else {
					log.Printf("[DEBUG] RewardToken合约连接成功，代币符号: %s", symbol)
				}

				// 检查调用账户是否有铸币权限
				isMinter, err := h.contractManager.RewardToken.AuthorizedMinters(nil, auth.From)
				if err != nil {
					log.Printf("[ERROR] 检查铸币权限失败: %v", err)
				} else if !isMinter {
					log.Printf("[WARN] 当前账户 %s 可能没有铸币权限", auth.From.Hex())
				} else {
					log.Printf("[DEBUG] 当前账户有铸币权限")
				}

				// 执行铸币操作
				tx, err := h.contractManager.RewardToken.Mint(auth, childAddress, tokenAmountInt)
				if err != nil {
					log.Printf("[ERROR] 铸造代币失败：%v", err)
				} else {
					log.Printf("[SUCCESS] 成功铸造 %s 代币给孩子（地址：%s），交易哈希：%s",
						tokenAmountInt.String(), task.AssignedChild.WalletAddress, tx.Hash().Hex())

					// 等待交易确认
					log.Printf("[DEBUG] 等待交易确认...")
					receipt, err := h.contractManager.WaitForTxReceipt(ctx, tx.Hash())
					if err != nil {
						log.Printf("[ERROR] 等待交易确认时出错: %v", err)
					} else {
						log.Printf("[SUCCESS] 交易已确认，区块号: %d，状态: %d",
							receipt.BlockNumber.Uint64(), receipt.Status)
					}
				}
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    task,
		"message": "Task approved and rewards transferred successfully",
	})
}

// RejectTask 拒绝任务
func (h *TaskHandler) RejectTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid task ID",
		})
		return
	}

	var req RejectTaskRequest
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

	role, exists := c.Get("role")
	if !exists || role != "parent" {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "Only parents can reject tasks",
		})
		return
	}

	// 查找任务
	var task models.Task
	result := h.db.First(&task, uint(id))
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Task not found",
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
	if task.CreatedBy != walletAddress.(string) {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "Only task creator can reject the task",
		})
		return
	}

	// 检查任务状态
	if task.Status != "completed" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Task is not completed yet",
		})
		return
	}

	// 开始事务
	tx := h.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 更新任务状态
	now := time.Now()
	task.Status = "rejected"
	task.RejectedAt = &now
	if req.Reason != "" {
		task.RejectionReason = &req.Reason
	}

	if err := tx.Save(&task).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to reject task",
		})
		return
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to commit transaction",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    task,
		"message": "Task rejected and reward refunded to parent",
	})
}

// DirectAssignTask 直接分配任务，跳过区块链调用
func (h *TaskHandler) DirectAssignTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid task ID",
		})
		return
	}

	var req struct {
		AssignedChildID string `json:"assigned_child_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request data",
		})
		return
	}

	// 转换childId为uint
	assignedChildID, err := strconv.ParseUint(req.AssignedChildID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid child ID format",
		})
		return
	}
	childID := uint(assignedChildID)

	walletAddress, exists := c.Get("wallet_address")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "User not authenticated",
		})
		return
	}

	// 查找任务
	var task models.Task
	result := h.db.First(&task, uint(id))
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Task not found",
		})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Database error",
		})
		return
	}

	// 检查权限（只有任务创建者可以更新）
	if task.CreatedBy != walletAddress.(string) {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "Only task creator can update the task",
		})
		return
	}

	// 验证孩子是否属于当前父母
	var child models.Child
	result = h.db.Where("id = ? AND parent_address = ?", childID, walletAddress).First(&child)
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Child not found or not belongs to you",
		})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Database error",
		})
		return
	}

	// 更新任务分配
	task.AssignedChildID = &childID

	// 如果分配给了孩子且状态是pending，改为in_progress
	if task.Status == "pending" {
		task.Status = "in_progress"
	}

	// 保存到数据库
	if err := h.db.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to update task",
		})
		return
	}

	// 预加载关联数据
	h.db.Preload("AssignedChild").First(&task, task.ID)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    task,
	})
}

// initializeRewardToken 尝试手动初始化奖励代币合约
func (h *TaskHandler) initializeRewardToken(tokenAddress string) error {
	if h.contractManager == nil {
		return fmt.Errorf("合约管理器未初始化")
	}

	// 检查是否已经初始化
	if h.contractManager.RewardToken != nil {
		return nil // 已经初始化，不需要再次初始化
	}

	// 使用新的公共方法初始化代币合约
	return h.contractManager.InitRewardToken(tokenAddress)
}
