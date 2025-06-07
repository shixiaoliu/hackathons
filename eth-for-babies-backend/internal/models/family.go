package models

import (
	"time"

	"gorm.io/gorm"
)

type Family struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Name          string         `json:"name" gorm:"not null"`
	ParentAddress string         `json:"parent_address" gorm:"uniqueIndex;not null"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`

	// 关联关系
	Parent   *User   `json:"parent,omitempty" gorm:"foreignKey:ParentAddress;references:WalletAddress"`
	Children []Child `json:"children,omitempty" gorm:"foreignKey:ParentAddress;references:ParentAddress"`
}

func (Family) TableName() string {
	return "families"
}