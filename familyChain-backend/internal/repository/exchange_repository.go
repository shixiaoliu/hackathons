package repository

import (
	"fmt"
	"time"

	"eth-for-babies-backend/internal/models"

	"gorm.io/gorm"
)

// ExchangeRepository 定义了兑换记录数据库操作
type ExchangeRepository struct {
	db *gorm.DB
}

// NewExchangeRepository 创建一个新的ExchangeRepository实例
func NewExchangeRepository(db *gorm.DB) *ExchangeRepository {
	return &ExchangeRepository{
		db: db,
	}
}

// Create 创建一个新的兑换记录
func (r *ExchangeRepository) Create(exchange *models.Exchange) error {
	return r.db.Create(exchange).Error
}

// GetByID 根据ID获取兑换记录
func (r *ExchangeRepository) GetByID(id uint) (*models.Exchange, error) {
	var exchange models.Exchange
	err := r.db.First(&exchange, id).Error
	if err != nil {
		return nil, err
	}
	return &exchange, nil
}

// GetByChildID 获取孩子的兑换记录
func (r *ExchangeRepository) GetByChildID(childID uint) ([]*models.Exchange, error) {
	var exchanges []*models.Exchange
	err := r.db.Where("child_id = ?", childID).
		Order("exchange_date DESC").
		Find(&exchanges).Error

	return exchanges, err
}

// GetByRewardID 获取奖品的兑换记录
func (r *ExchangeRepository) GetByRewardID(rewardID uint) ([]*models.Exchange, error) {
	var exchanges []*models.Exchange
	err := r.db.Where("reward_id = ?", rewardID).
		Order("exchange_date DESC").
		Find(&exchanges).Error

	return exchanges, err
}

// GetByFamilyID 获取家庭的兑换记录
func (r *ExchangeRepository) GetByFamilyID(familyID uint) ([]*models.Exchange, error) {
	var exchanges []*models.Exchange
	err := r.db.Table("exchanges").
		Joins("JOIN children ON exchanges.child_id = children.id").
		Where("children.family_id = ?", familyID).
		Order("exchanges.exchange_date DESC").
		Find(&exchanges).Error

	return exchanges, err
}

// UpdateStatus 更新兑换记录状态
func (r *ExchangeRepository) UpdateStatus(id uint, status models.ExchangeStatus, notes string) error {
	updates := map[string]interface{}{
		"status":     status,
		"notes":      notes,
		"updated_at": time.Now(),
	}

	// 如果状态是已完成，更新完成日期
	if status == models.ExchangeStatusCompleted {
		updates["completed_date"] = time.Now()
	}

	return r.db.Model(&models.Exchange{}).Where("id = ?", id).Updates(updates).Error
}

// Delete 删除兑换记录
func (r *ExchangeRepository) Delete(id uint) error {
	return r.db.Delete(&models.Exchange{}, id).Error
}

// WithTransaction 在事务中执行操作
func (r *ExchangeRepository) WithTransaction(fn func(*ExchangeRepository) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		txRepo := &ExchangeRepository{db: tx}
		return fn(txRepo)
	})
}

// GetExchangeWithDetails 获取带详细信息的兑换记录
func (r *ExchangeRepository) GetExchangeWithDetails(id uint) (*models.Exchange, error) {
	var exchange models.Exchange

	err := r.db.First(&exchange, id).Error
	if err != nil {
		return nil, err
	}

	// 获取奖品名称和图片
	var reward models.Reward
	if err := r.db.Select("name, image_url").First(&reward, exchange.RewardID).Error; err == nil {
		exchange.RewardName = reward.Name
		exchange.RewardImage = reward.ImageURL
	}

	// 获取孩子名称
	var child models.Child
	if err := r.db.Select("name").First(&child, exchange.ChildID).Error; err == nil {
		exchange.ChildName = child.Name
	}

	return &exchange, nil
}

// GetChildExchangesWithDetails 获取孩子的带详细信息的兑换记录
func (r *ExchangeRepository) GetChildExchangesWithDetails(childID uint) ([]*models.Exchange, error) {
	var exchanges []*models.Exchange

	fmt.Printf("获取孩子兑换记录，childID: %d\n", childID)

	err := r.db.Where("child_id = ?", childID).
		Order("exchange_date DESC").
		Find(&exchanges).Error

	if err != nil {
		fmt.Printf("获取孩子兑换记录失败: %v\n", err)
		return nil, err
	}

	fmt.Printf("查询到 %d 条兑换记录\n", len(exchanges))

	for _, exchange := range exchanges {
		// 获取奖品名称和图片
		var reward models.Reward
		if err := r.db.Select("name, image_url").First(&reward, exchange.RewardID).Error; err == nil {
			exchange.RewardName = reward.Name
			exchange.RewardImage = reward.ImageURL
			fmt.Printf("兑换记录 %d: 奖品名称=%s, 奖品图片=%s\n", exchange.ID, reward.Name, reward.ImageURL)
		} else {
			fmt.Printf("获取奖品信息失败，兑换ID: %d, 奖品ID: %d, 错误: %v\n", exchange.ID, exchange.RewardID, err)
		}
	}

	return exchanges, nil
}
