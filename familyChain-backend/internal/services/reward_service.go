package services

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"eth-for-babies-backend/internal/models"
	"eth-for-babies-backend/internal/repository"
	"eth-for-babies-backend/pkg/blockchain"
)

// RewardService 处理奖品和兑换相关的业务逻辑
type RewardService struct {
	rewardRepo     *repository.RewardRepository
	exchangeRepo   *repository.ExchangeRepository
	childRepo      *repository.ChildRepository
	contractClient *blockchain.ContractManager
}

// NewRewardService 创建一个新的奖励服务
func NewRewardService(
	rewardRepo *repository.RewardRepository,
	exchangeRepo *repository.ExchangeRepository,
	childRepo *repository.ChildRepository,
	contractClient *blockchain.ContractManager,
) *RewardService {
	return &RewardService{
		rewardRepo:     rewardRepo,
		exchangeRepo:   exchangeRepo,
		childRepo:      childRepo,
		contractClient: contractClient,
	}
}

// CreateReward 创建新的实物奖励
func (s *RewardService) CreateReward(ctx context.Context, userID uint, familyID uint, req models.RewardCreateRequest) (uint, error) {
	// 验证家庭是否存在
	// TODO: 检查用户是否是该家庭的家长

	// 添加详细日志
	fmt.Printf("开始创建奖品 - 用户ID: %d, 家庭ID: %d, 奖品名称: %s\n", userID, familyID, req.Name)
	fmt.Printf("请求数据详情: %+v\n", req)

	// 创建数据库记录
	reward := &models.Reward{
		FamilyID:    familyID,
		Name:        req.Name,
		Description: req.Description,
		ImageURL:    req.ImageURL,
		TokenPrice:  req.TokenPrice,
		CreatedBy:   userID,
		Active:      true,
		Stock:       req.Stock,
	}

	fmt.Printf("准备保存到数据库的奖品记录: %+v\n", reward)

	if err := s.rewardRepo.Create(reward); err != nil {
		fmt.Printf("数据库创建奖品记录失败: %v\n", err)
		return 0, fmt.Errorf("failed to create reward in database: %w", err)
	}

	fmt.Printf("奖品创建成功 - ID: %d\n", reward.ID)

	// 返回奖品ID
	return reward.ID, nil
}

// GetReward 获取奖品详情
func (s *RewardService) GetReward(ctx context.Context, id uint) (*models.Reward, error) {
	return s.rewardRepo.GetByID(id)
}

// GetFamilyRewards 获取家庭的所有奖品
func (s *RewardService) GetFamilyRewards(ctx context.Context, familyID uint, activeOnly bool) ([]*models.Reward, error) {
	return s.rewardRepo.GetByFamilyID(familyID, activeOnly)
}

// UpdateReward 更新奖品信息
func (s *RewardService) UpdateReward(ctx context.Context, id uint, req models.RewardUpdateRequest) error {
	// 检查奖品是否存在
	reward, err := s.rewardRepo.GetByID(id)
	if err != nil {
		fmt.Printf("获取奖品失败: %v\n", err)
		return fmt.Errorf("failed to get reward: %w", err)
	}
	if reward == nil {
		fmt.Println("奖品不存在")
		return fmt.Errorf("reward not found")
	}

	fmt.Printf("更新奖品ID: %d, 请求数据: %+v\n", id, req)

	// TODO: 检查用户是否有权限更新奖品

	// 临时跳过区块链调用
	/*
		// 更新链上奖品信息
		rewardIdBig := new(big.Int).SetUint64(uint64(id))
		tokenPriceBig := new(big.Int).SetInt64(int64(req.TokenPrice))
		stockBig := new(big.Int).SetInt64(int64(req.Stock))
		active := false
		if req.Active != nil {
			active = *req.Active
		}

		txOpts, err := s.contractClient.GetTransactOpts(ctx)
		if err != nil {
			return fmt.Errorf("failed to get transaction options: %w", err)
		}

		tx, err := s.contractClient.RewardRegistry.UpdateReward(
			txOpts,
			rewardIdBig,
			req.Name,
			req.Description,
			req.ImageURL,
			tokenPriceBig,
			stockBig,
			active,
		)
		if err != nil {
			return fmt.Errorf("failed to update reward on chain: %w", err)
		}

		// 等待交易确认
		_, err = s.contractClient.WaitForTxReceipt(ctx, tx.Hash())
		if err != nil {
			return fmt.Errorf("transaction failed: %w", err)
		}
	*/

	// 更新数据库记录
	updates := make(map[string]interface{})
	if req.Name != nil {
		updates["name"] = *req.Name
	}
	if req.Description != nil {
		updates["description"] = *req.Description
	}
	if req.ImageURL != nil {
		updates["image_url"] = *req.ImageURL
	}
	if req.TokenPrice != nil && *req.TokenPrice > 0 {
		updates["token_price"] = *req.TokenPrice
	}
	if req.Stock != nil && *req.Stock >= 0 {
		updates["stock"] = *req.Stock
	}
	if req.Active != nil {
		updates["active"] = *req.Active
	}

	fmt.Printf("更新数据库记录: %+v\n", updates)
	err = s.rewardRepo.Update(id, updates)
	if err != nil {
		fmt.Printf("更新奖品数据库记录失败: %v\n", err)
		return fmt.Errorf("failed to update reward in database: %w", err)
	}
	fmt.Printf("奖品更新成功, ID: %d\n", id)

	return nil
}

// ExchangeReward 兑换奖品
func (s *RewardService) ExchangeReward(ctx context.Context, childID uint, req models.ExchangeCreateRequest) (uint, error) {
	// 获取孩子信息
	child, err := s.childRepo.GetByID(childID)
	if err != nil {
		return 0, fmt.Errorf("failed to get child: %w", err)
	}
	if child == nil {
		return 0, fmt.Errorf("child not found")
	}

	// 获取奖品信息
	reward, err := s.rewardRepo.GetByID(uint(req.RewardID))
	if err != nil {
		return 0, fmt.Errorf("failed to get reward: %w", err)
	}
	if reward == nil {
		return 0, fmt.Errorf("reward not found")
	}

	// 检查奖品是否可用
	if !reward.Active {
		return 0, fmt.Errorf("reward is not active")
	}
	if reward.Stock <= 0 {
		return 0, fmt.Errorf("reward is out of stock")
	}

	// 获取孩子的以太坊地址
	childAddress := common.HexToAddress(child.WalletAddress)
	if (childAddress == common.Address{}) {
		return 0, fmt.Errorf("invalid child wallet address")
	}

	// 使用孩子的地址调用合约进行兑换
	rewardIdBig := new(big.Int).SetUint64(uint64(req.RewardID))

	// 创建数据库兑换记录 - 先创建记录，即使区块链交易失败也能跟踪
	exchange := &models.Exchange{
		RewardID:    uint(req.RewardID),
		ChildID:     childID,
		TokenAmount: reward.TokenPrice,
		Status:      models.ExchangeStatusPending,
		Notes:       req.Notes,
	}

	if err := s.exchangeRepo.Create(exchange); err != nil {
		return 0, fmt.Errorf("failed to create exchange in database: %w", err)
	}

	// 启动一个goroutine来处理区块链交易，避免阻塞API响应
	go func() {
		// 创建新的上下文，不受原始请求上下文的限制
		bgCtx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
		defer cancel()

		fmt.Printf("开始后台处理兑换交易，兑换ID: %d, 奖品ID: %d, 前端已销毁代币: %v\n",
			exchange.ID, req.RewardID, req.TokenBurned)

		// 无论前端是否已经销毁了代币，都调用区块链合约进行兑换
		// 这确保了代币被正确销毁并记录在区块链上
		fmt.Printf("调用区块链合约进行兑换和代币销毁\n")

		// 改进的重试机制
		var tx *types.Transaction
		var txErr error
		maxRetries := 5
		backoffTime := 2 * time.Second // 初始退避时间

		for attempt := 0; attempt < maxRetries; attempt++ {
			// 获取孩子的交易选项
			txCtx, txCancel := context.WithTimeout(bgCtx, 15*time.Second)
			txOpts, err := s.contractClient.GetChildTransactOpts(txCtx, childAddress)
			txCancel()

			if err != nil {
				fmt.Printf("获取交易选项失败 (尝试 %d/%d): %v\n", attempt+1, maxRetries, err)
				time.Sleep(backoffTime)
				backoffTime *= 2 // 指数退避
				continue
			}

			// 记录尝试信息
			fmt.Printf("尝试兑换奖品 (尝试 %d/%d), 使用nonce: %d\n",
				attempt+1, maxRetries, txOpts.Nonce.Uint64())

			// 执行交易
			txCtx, txCancel = context.WithTimeout(bgCtx, 20*time.Second)
			tx, txErr = s.contractClient.RewardRegistry.ExchangeReward(txOpts, rewardIdBig)
			txCancel()

			// 如果成功，跳出循环
			if txErr == nil {
				fmt.Printf("交易提交成功，交易哈希: %s\n", tx.Hash().Hex())
				break
			}

			// 详细记录错误信息
			fmt.Printf("交易失败 (尝试 %d/%d): %v\n", attempt+1, maxRetries, txErr)

			// 分析错误类型
			if strings.Contains(txErr.Error(), "nonce too low") ||
				strings.Contains(txErr.Error(), "replacement transaction underpriced") {
				fmt.Printf("检测到nonce或价格错误，将在下次尝试时使用更新的值\n")
				time.Sleep(backoffTime)
				backoffTime *= 2
				continue
			} else if strings.Contains(txErr.Error(), "insufficient funds") {
				// 更新兑换状态为失败
				updateErr := s.exchangeRepo.UpdateStatus(exchange.ID, models.ExchangeStatusFailed, "账户余额不足")
				if updateErr != nil {
					fmt.Printf("更新兑换状态失败: %v\n", updateErr)
				}
				return // 终止goroutine
			}

			// 其他错误，等待后重试
			time.Sleep(backoffTime)
			backoffTime *= 2
		}

		// 检查是否所有尝试都失败了
		if txErr != nil {
			fmt.Printf("所有交易尝试均失败，更新兑换状态为失败\n")
			updateErr := s.exchangeRepo.UpdateStatus(exchange.ID, models.ExchangeStatusFailed, "区块链交易错误，请稍后再试")
			if updateErr != nil {
				fmt.Printf("更新兑换状态失败: %v\n", updateErr)
			}
			return
		}

		// 交易已提交，但不等待确认，避免阻塞
		fmt.Printf("交易已提交，哈希: %s，更新兑换状态为处理中\n", tx.Hash().Hex())

		// 更新奖品库存
		if err := s.rewardRepo.UpdateStock(uint(req.RewardID), -1); err != nil {
			fmt.Printf("更新奖品库存失败: %v\n", err)
		}

		// 启动另一个goroutine来等待交易确认
		go func() {
			// 创建新的上下文，专门用于等待交易确认
			confirmCtx, confirmCancel := context.WithTimeout(context.Background(), 5*time.Minute)
			defer confirmCancel()

			fmt.Printf("开始等待交易确认，交易哈希: %s\n", tx.Hash().Hex())
			receipt, err := s.contractClient.WaitForTxReceipt(confirmCtx, tx.Hash())

			if err != nil {
				fmt.Printf("交易确认失败: %v\n", err)
				// 不更新状态，保持为处理中
				return
			}

			// 检查交易状态
			if receipt.Status == types.ReceiptStatusSuccessful {
				fmt.Printf("交易确认成功，区块号: %d\n", receipt.BlockNumber.Uint64())
				// 更新兑换状态为已确认
				updateErr := s.exchangeRepo.UpdateStatus(exchange.ID, models.ExchangeStatusConfirmed, "区块链交易已确认")
				if updateErr != nil {
					fmt.Printf("更新兑换状态失败: %v\n", updateErr)
				}
			} else {
				fmt.Printf("交易失败，状态码: %d\n", receipt.Status)
				// 更新兑换状态为失败
				updateErr := s.exchangeRepo.UpdateStatus(exchange.ID, models.ExchangeStatusFailed, "区块链交易执行失败")
				if updateErr != nil {
					fmt.Printf("更新兑换状态失败: %v\n", updateErr)
				}

				// 恢复奖品库存
				if err := s.rewardRepo.UpdateStock(uint(req.RewardID), 1); err != nil {
					fmt.Printf("恢复奖品库存失败: %v\n", err)
				}
			}
		}()
	}()

	// 立即返回兑换ID，不等待区块链交易完成
	return exchange.ID, nil
}

// UpdateExchangeStatus 更新兑换状态
func (s *RewardService) UpdateExchangeStatus(ctx context.Context, exchangeID uint, req models.ExchangeUpdateRequest) error {
	// 获取兑换记录
	exchange, err := s.exchangeRepo.GetByID(exchangeID)
	if err != nil {
		return fmt.Errorf("failed to get exchange: %w", err)
	}
	if exchange == nil {
		return fmt.Errorf("exchange not found")
	}

	// 获取奖品信息
	reward, err := s.rewardRepo.GetByID(exchange.RewardID)
	if err != nil {
		return fmt.Errorf("failed to get reward: %w", err)
	}
	if reward == nil {
		return fmt.Errorf("reward not found")
	}

	// TODO: 检查用户是否有权限更新兑换状态

	// 如果状态是已完成，更新链上兑换状态
	if req.Status == models.ExchangeStatusCompleted {
		exchangeIdBig := new(big.Int).SetUint64(uint64(exchangeID))

		txOpts, err := s.contractClient.GetTransactOpts(ctx)
		if err != nil {
			return fmt.Errorf("failed to get transaction options: %w", err)
		}

		tx, err := s.contractClient.RewardRegistry.FulfillExchange(txOpts, exchangeIdBig)
		if err != nil {
			return fmt.Errorf("failed to fulfill exchange on chain: %w", err)
		}

		// 等待交易确认
		_, err = s.contractClient.WaitForTxReceipt(ctx, tx.Hash())
		if err != nil {
			return fmt.Errorf("transaction failed: %w", err)
		}
	}

	// 更新数据库状态
	return s.exchangeRepo.UpdateStatus(exchangeID, req.Status, req.Notes)
}

// GetChildExchanges 获取孩子的兑换记录
func (s *RewardService) GetChildExchanges(ctx context.Context, childID uint) ([]*models.Exchange, error) {
	return s.exchangeRepo.GetChildExchangesWithDetails(childID)
}

// GetFamilyExchanges 获取家庭的兑换记录
func (s *RewardService) GetFamilyExchanges(ctx context.Context, familyID uint) ([]*models.Exchange, error) {
	return s.exchangeRepo.GetByFamilyID(familyID)
}

// GetExchange 获取兑换详情
func (s *RewardService) GetExchange(ctx context.Context, exchangeID uint) (*models.Exchange, error) {
	return s.exchangeRepo.GetExchangeWithDetails(exchangeID)
}
