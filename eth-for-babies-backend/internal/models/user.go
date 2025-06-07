package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	WalletAddress string         `json:"wallet_address" gorm:"uniqueIndex;not null"`
	Role          string         `json:"role" gorm:"not null;check:role IN ('parent', 'child', 'temp')"`
	Nonce         string         `json:"-" gorm:"not null"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`

	// 关联关系
	Family   *Family  `json:"family,omitempty" gorm:"foreignKey:ParentAddress;references:WalletAddress"`
	Children []Child  `json:"children,omitempty" gorm:"foreignKey:ParentAddress;references:WalletAddress"`
	Tasks    []Task   `json:"tasks,omitempty" gorm:"foreignKey:CreatedBy;references:WalletAddress"`
}

func (User) TableName() string {
	return "users"
}