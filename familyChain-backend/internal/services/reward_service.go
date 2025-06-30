package services

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

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

	// 获取孩子的交易选项
	txOpts, err := s.contractClient.GetChildTransactOpts(ctx, childAddress)
	if err != nil {
		return 0, fmt.Errorf("failed to get transaction options: %w", err)
	}

	tx, err := s.contractClient.RewardRegistry.ExchangeReward(txOpts, rewardIdBig)
	if err != nil {
		return 0, fmt.Errorf("failed to exchange reward on chain: %w", err)
	}

	// 等待交易确认
	_, err = s.contractClient.WaitForTxReceipt(ctx, tx.Hash())
	if err != nil {
		return 0, fmt.Errorf("transaction failed: %w", err)
	}

	// 获取交换ID（从事件中解析）
	// 简化处理，这里使用交换计数作为ID
	callOpts := &bind.CallOpts{Context: ctx}
	exchangeCount, err := s.contractClient.RewardRegistry.ExchangeCount(callOpts)
	if err != nil {
		return 0, fmt.Errorf("failed to get exchange count: %w", err)
	}

	// 创建数据库兑换记录
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

	// 更新奖品库存
	if err := s.rewardRepo.UpdateStock(uint(req.RewardID), -1); err != nil {
		return 0, fmt.Errorf("failed to update reward stock: %w", err)
	}

	return uint(exchangeCount.Uint64()), nil
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
