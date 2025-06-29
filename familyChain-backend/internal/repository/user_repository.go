package repository

import (
	"eth-for-babies-backend/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create 创建用户
func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// GetByID 根据ID获取用户
func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByWalletAddress 根据钱包地址获取用户
func (r *UserRepository) GetByWalletAddress(walletAddress string) (*models.User, error) {
	var user models.User
	err := r.db.Where("wallet_address = ?", walletAddress).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update 更新用户
func (r *UserRepository) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&models.User{}).Where("id = ?", id).Updates(updates).Error
}

// Delete 删除用户
func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

// UpdateNonce 更新用户nonce
func (r *UserRepository) UpdateNonce(walletAddress string, nonce string) error {
	return r.db.Model(&models.User{}).Where("wallet_address = ?", walletAddress).Update("nonce", nonce).Error
}

// GetByRole 根据角色获取用户列表
func (r *UserRepository) GetByRole(role string) ([]*models.User, error) {
	var users []*models.User
	err := r.db.Where("role = ?", role).Find(&users).Error
	return users, err
}

// List 获取用户列表
func (r *UserRepository) List(limit, offset int) ([]*models.User, error) {
	var users []*models.User
	query := r.db.Order("created_at DESC")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	err := query.Find(&users).Error
	return users, err
}

// Count 获取用户总数
func (r *UserRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&models.User{}).Count(&count).Error
	return count, err
}

// CountByRole 根据角色获取用户数量
func (r *UserRepository) CountByRole(role string) (int64, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("role = ?", role).Count(&count).Error
	return count, err
}