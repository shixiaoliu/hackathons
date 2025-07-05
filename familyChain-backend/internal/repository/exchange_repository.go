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

	fmt.Printf("获取家庭兑换记录，familyID: %d\n", familyID)

	// 先获取该家庭的 parent_address
	var parentAddress string
	fmt.Printf("执行SQL: SELECT parent_address FROM families WHERE id = %d\n", familyID)
	err := r.db.Table("families").
		Select("parent_address").
		Where("id = ?", familyID).
		Pluck("parent_address", &parentAddress).Error

	if err != nil {
		fmt.Printf("获取家庭parent_address失败: %v\n", err)
		return nil, err
	}

	if parentAddress == "" {
		fmt.Printf("未找到家庭 %d 的parent_address\n", familyID)
		return []*models.Exchange{}, nil
	}

	fmt.Printf("家庭 %d 的parent_address: %s\n", familyID, parentAddress)

	// 获取该家庭下所有孩子的ID
	var childIDs []uint
	fmt.Printf("执行SQL: SELECT id FROM children WHERE parent_address = '%s'\n", parentAddress)
	err = r.db.Table("children").
		Select("id").
		Where("parent_address = ?", parentAddress).
		Pluck("id", &childIDs).Error

	if err != nil {
		fmt.Printf("获取家庭孩子ID失败: %v\n", err)
		return nil, err
	}

	fmt.Printf("家庭 %d 有 %d 个孩子，IDs: %v\n", familyID, len(childIDs), childIDs)

	if len(childIDs) == 0 {
		// 如果没有孩子，返回空数组
		fmt.Printf("家庭 %d 没有孩子，返回空兑换记录\n", familyID)
		return []*models.Exchange{}, nil
	}

	// 获取这些孩子的所有兑换记录
	fmt.Printf("执行SQL: SELECT * FROM exchanges WHERE child_id IN (%v) ORDER BY exchange_date DESC\n", childIDs)
	err = r.db.Where("child_id IN ?", childIDs).
		Order("exchange_date DESC").
		Find(&exchanges).Error

	if err != nil {
		fmt.Printf("获取家庭兑换记录失败: %v\n", err)
		return nil, err
	}

	fmt.Printf("查询到 %d 条兑换记录\n", len(exchanges))

	// 为每条记录添加详细信息
	for _, exchange := range exchanges {
		// 获取奖品名称和图片
		var reward models.Reward
		fmt.Printf("获取奖品信息，奖品ID: %d\n", exchange.RewardID)
		if err := r.db.Select("name, image_url").First(&reward, exchange.RewardID).Error; err == nil {
			exchange.RewardName = reward.Name
			exchange.RewardImage = reward.ImageURL
			fmt.Printf("兑换记录 %d: 奖品名称=%s\n", exchange.ID, reward.Name)
		} else {
			fmt.Printf("获取奖品信息失败，兑换ID: %d, 奖品ID: %d, 错误: %v\n", exchange.ID, exchange.RewardID, err)
		}

		// 获取孩子名称
		var child models.Child
		fmt.Printf("获取孩子信息，孩子ID: %d\n", exchange.ChildID)
		if err := r.db.Select("name").First(&child, exchange.ChildID).Error; err == nil {
			exchange.ChildName = child.Name
			fmt.Printf("兑换记录 %d: 孩子名称=%s\n", exchange.ID, child.Name)
		} else {
			fmt.Printf("获取孩子信息失败，兑换ID: %d, 孩子ID: %d, 错误: %v\n", exchange.ID, exchange.ChildID, err)
		}
	}

	fmt.Printf("返回家庭 %d 的 %d 条兑换记录\n", familyID, len(exchanges))
	return exchanges, nil
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
