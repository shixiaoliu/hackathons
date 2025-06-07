package repository

import (
	"eth-for-babies-backend/internal/models"
	"gorm.io/gorm"
)

type FamilyRepository struct {
	db *gorm.DB
}

func NewFamilyRepository(db *gorm.DB) *FamilyRepository {
	return &FamilyRepository{db: db}
}

// Create 创建家庭
func (r *FamilyRepository) Create(family *models.Family) error {
	return r.db.Create(family).Error
}

// GetByID 根据ID获取家庭
func (r *FamilyRepository) GetByID(id uint) (*models.Family, error) {
	var family models.Family
	err := r.db.Preload("Parent").Preload("Children").First(&family, id).Error
	if err != nil {
		return nil, err
	}
	return &family, nil
}

// GetByParentAddress 根据家长地址获取家庭
func (r *FamilyRepository) GetByParentAddress(parentAddress string) (*models.Family, error) {
	var family models.Family
	err := r.db.Preload("Parent").Preload("Children").Where("parent_address = ?", parentAddress).First(&family).Error
	if err != nil {
		return nil, err
	}
	return &family, nil
}

// Update 更新家庭
func (r *FamilyRepository) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&models.Family{}).Where("id = ?", id).Updates(updates).Error
}

// Delete 删除家庭
func (r *FamilyRepository) Delete(id uint) error {
	return r.db.Delete(&models.Family{}, id).Error
}

// List 获取家庭列表
func (r *FamilyRepository) List(limit, offset int) ([]*models.Family, error) {
	var families []*models.Family
	query := r.db.Preload("Parent").Preload("Children").Order("created_at DESC")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	err := query.Find(&families).Error
	return families, err
}

// GetFamiliesWithChildren 获取包含孩子信息的家庭列表
func (r *FamilyRepository) GetFamiliesWithChildren() ([]*models.Family, error) {
	var families []*models.Family
	err := r.db.Preload("Parent").Preload("Children").Find(&families).Error
	return families, err
}

// Count 获取家庭总数
func (r *FamilyRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&models.Family{}).Count(&count).Error
	return count, err
}

// GetFamilyStatistics 获取家庭统计信息
func (r *FamilyRepository) GetFamilyStatistics(id uint) (map[string]interface{}, error) {
	family, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}

	// 统计孩子数量
	var childCount int64
	r.db.Model(&models.Child{}).Where("family_id = ?", id).Count(&childCount)

	// 统计任务数量
	var taskCount int64
	r.db.Model(&models.Task{}).Where("created_by = ?", family.ParentAddress).Count(&taskCount)

	// 统计完成的任务数量
	var completedTaskCount int64
	r.db.Model(&models.Task{}).Where("created_by = ? AND status = ?", family.ParentAddress, "approved").Count(&completedTaskCount)

	// 计算总奖励
	var totalRewards float64
	r.db.Model(&models.Child{}).Where("family_id = ?", id).Select("COALESCE(SUM(total_rewards), 0)").Scan(&totalRewards)

	stats := map[string]interface{}{
		"family_id":            id,
		"family_name":          family.Name,
		"children_count":       childCount,
		"total_tasks":          taskCount,
		"completed_tasks":      completedTaskCount,
		"total_rewards":        totalRewards,
		"created_at":           family.CreatedAt,
	}

	return stats, nil
}

// WithTransaction 在事务中执行操作
func (r *FamilyRepository) WithTransaction(fn func(*FamilyRepository) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		txRepo := &FamilyRepository{db: tx}
		return fn(txRepo)
	})
}