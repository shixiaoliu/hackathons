package repository

import (
	"time"

	"eth-for-babies-backend/internal/models"

	"gorm.io/gorm"
)

// RewardRepository 定义了奖品数据库操作的接口
type RewardRepository struct {
	db *gorm.DB
}

// NewRewardRepository 创建一个新的RewardRepository实例
func NewRewardRepository(db *gorm.DB) *RewardRepository {
	return &RewardRepository{
		db: db,
	}
}

// Create 创建一个新的奖品记录
func (r *RewardRepository) Create(reward *models.Reward) error {
	return r.db.Create(reward).Error
}

// GetByID 根据ID获取奖品
func (r *RewardRepository) GetByID(id uint) (*models.Reward, error) {
	var reward models.Reward
	err := r.db.First(&reward, id).Error
	if err != nil {
		return nil, err
	}
	return &reward, nil
}

// GetByFamilyID 根据家庭ID获取奖品列表
func (r *RewardRepository) GetByFamilyID(familyID uint, activeOnly bool) ([]*models.Reward, error) {
	var rewards []*models.Reward
	query := r.db.Where("family_id = ?", familyID)

	if activeOnly {
		query = query.Where("active = ?", true)
	}

	err := query.Order("created_at DESC").Find(&rewards).Error
	return rewards, err
}

// Update 更新奖品信息
func (r *RewardRepository) Update(id uint, updates map[string]interface{}) error {
	updates["updated_at"] = time.Now()
	return r.db.Model(&models.Reward{}).Where("id = ?", id).Updates(updates).Error
}

// Delete 删除奖品
func (r *RewardRepository) Delete(id uint) error {
	return r.db.Delete(&models.Reward{}, id).Error
}

// UpdateStock 更新奖品库存
func (r *RewardRepository) UpdateStock(id uint, stockChange int) error {
	var reward models.Reward

	// 使用事务确保库存更新的原子性
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 查询当前库存
		if err := tx.First(&reward, id).Error; err != nil {
			return err
		}

		// 计算新库存
		newStock := reward.Stock + stockChange

		// 检查库存是否足够
		if newStock < 0 {
			return gorm.ErrInvalidData
		}

		// 更新库存
		return tx.Model(&models.Reward{}).Where("id = ?", id).Updates(map[string]interface{}{
			"stock":      newStock,
			"updated_at": time.Now(),
		}).Error
	})
}

// WithTransaction 在事务中执行操作
func (r *RewardRepository) WithTransaction(fn func(*RewardRepository) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		txRepo := &RewardRepository{db: tx}
		return fn(txRepo)
	})
}
