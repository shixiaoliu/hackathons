package repository

import (
	"gorm.io/gorm"
)

// BaseRepository 基础仓库接口
type BaseRepository interface {
	Create(entity interface{}) error
	GetByID(id uint, entity interface{}) error
	Update(id uint, updates map[string]interface{}) error
	Delete(id uint, entity interface{}) error
	List(entities interface{}, conditions ...interface{}) error
}

// baseRepository 基础仓库实现
type baseRepository struct {
	db *gorm.DB
}

// NewBaseRepository 创建基础仓库
func NewBaseRepository(db *gorm.DB) BaseRepository {
	return &baseRepository{db: db}
}

// Create 创建实体
func (r *baseRepository) Create(entity interface{}) error {
	return r.db.Create(entity).Error
}

// GetByID 根据ID获取实体
func (r *baseRepository) GetByID(id uint, entity interface{}) error {
	return r.db.First(entity, id).Error
}

// Update 更新实体
func (r *baseRepository) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&struct{ ID uint }{ID: id}).Updates(updates).Error
}

// Delete 删除实体
func (r *baseRepository) Delete(id uint, entity interface{}) error {
	return r.db.Delete(entity, id).Error
}

// List 列出实体
func (r *baseRepository) List(entities interface{}, conditions ...interface{}) error {
	query := r.db
	for i := 0; i < len(conditions); i += 2 {
		if i+1 < len(conditions) {
			query = query.Where(conditions[i], conditions[i+1])
		}
	}
	return query.Find(entities).Error
}