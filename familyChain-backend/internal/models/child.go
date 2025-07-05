package models

import (
	"time"

	"gorm.io/gorm"
)

type Child struct {
	ID                  uint           `json:"id" gorm:"primaryKey"`
	Name                string         `json:"name" gorm:"not null"`
	WalletAddress       string         `json:"wallet_address" gorm:"uniqueIndex;not null"`
	Age                 int            `json:"age" gorm:"not null"`
	Avatar              *string        `json:"avatar,omitempty"`
	ParentAddress       string         `json:"parent_address" gorm:"not null"`
	TotalTasksCompleted int            `json:"total_tasks_completed" gorm:"default:0"`
	TotalRewardsEarned  string         `json:"total_rewards_earned" gorm:"default:'0'"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `json:"-" gorm:"index"`

	// 关联关系
	Parent *User   `json:"parent,omitempty" gorm:"foreignKey:ParentAddress;references:WalletAddress"`
	Family *Family `json:"family,omitempty" gorm:"foreignKey:ParentAddress;references:ParentAddress"`
	Tasks  []Task  `json:"tasks,omitempty" gorm:"foreignKey:AssignedChildID;references:ID"`
}

func (Child) TableName() string {
	return "children"
}
