package repository

import (
	"eth-for-babies-backend/internal/models"

	"gorm.io/gorm"
)

type ChildRepository struct {
	db *gorm.DB
}

func NewChildRepository(db *gorm.DB) *ChildRepository {
	return &ChildRepository{db: db}
}

// Create 创建孩子
func (r *ChildRepository) Create(child *models.Child) error {
	return r.db.Create(child).Error
}

// GetByID 根据ID获取孩子
func (r *ChildRepository) GetByID(id uint) (*models.Child, error) {
	var child models.Child
	err := r.db.Preload("Parent").Preload("Family").Preload("Tasks").First(&child, id).Error
	if err != nil {
		return nil, err
	}
	return &child, nil
}

// GetByWalletAddress 根据钱包地址获取孩子
func (r *ChildRepository) GetByWalletAddress(walletAddress string) (*models.Child, error) {
	var child models.Child
	err := r.db.Preload("Parent").Preload("Family").Where("wallet_address = ?", walletAddress).First(&child).Error
	if err != nil {
		return nil, err
	}
	return &child, nil
}

// GetByParentAddress 根据家长地址获取孩子列表
func (r *ChildRepository) GetByParentAddress(parentAddress string) ([]*models.Child, error) {
	var children []*models.Child
	err := r.db.Preload("Parent").Preload("Family").Where("parent_address = ?", parentAddress).Find(&children).Error
	return children, err
}

// GetByFamilyID 根据家庭ID获取孩子列表
func (r *ChildRepository) GetByFamilyID(familyID uint) ([]*models.Child, error) {
	var children []*models.Child
	err := r.db.Preload("Parent").Preload("Family").Where("family_id = ?", familyID).Find(&children).Error
	return children, err
}

// Update 更新孩子
func (r *ChildRepository) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&models.Child{}).Where("id = ?", id).Updates(updates).Error
}

// DeleteChild 删除孩子
func (r *ChildRepository) DeleteChild(id uint) error {
	return r.db.Delete(&models.Child{}, id).Error
}

// UpdateWithRaw 使用原始SQL更新孩子（用于统计字段的增量更新）
func (r *ChildRepository) UpdateWithRaw(id uint, updates map[string]interface{}, args ...interface{}) error {
	// 对于包含原始SQL的更新，直接使用Updates方法
	// GORM会自动处理原始SQL表达式
	query := r.db.Model(&models.Child{}).Where("id = ?", id)

	// 如果有额外参数，需要特殊处理
	if len(args) > 0 {
		// 对于包含参数的原始SQL更新
		for field, value := range updates {
			if valueStr, ok := value.(string); ok {
				// 执行原始SQL更新
				return r.db.Model(&models.Child{}).Where("id = ?", id).Update(field, gorm.Expr(valueStr, args...)).Error
			}
		}
	}

	return query.Updates(updates).Error
}
