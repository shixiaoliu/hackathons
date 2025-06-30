package models

import "time"

// Reward 表示家长创建的实物奖励
type Reward struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	FamilyID    uint      `json:"family_id" gorm:"not null;index"`
	Name        string    `json:"name" gorm:"not null;size:255"`
	Description string    `json:"description" gorm:"type:text"`
	ImageURL    string    `json:"image_url" gorm:"type:text"`
	TokenPrice  int       `json:"token_price" gorm:"not null"`
	CreatedBy   uint      `json:"created_by" gorm:"not null"`
	Active      bool      `json:"active" gorm:"default:true;index"`
	Stock       int       `json:"stock" gorm:"default:1"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// RewardCreateRequest 表示创建奖品的请求
type RewardCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	TokenPrice  int    `json:"token_price" binding:"required,min=1"`
	Stock       int    `json:"stock" binding:"min=0"`
}

// RewardUpdateRequest 表示更新奖品的请求
type RewardUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	TokenPrice  int    `json:"token_price" binding:"min=1"`
	Active      *bool  `json:"active"`
	Stock       int    `json:"stock" binding:"min=0"`
}
