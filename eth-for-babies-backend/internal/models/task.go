package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	Title            string         `json:"title" gorm:"not null"`
	Description      string         `json:"description" gorm:"not null"`
	RewardAmount     string         `json:"reward_amount" gorm:"not null"`
	Difficulty       string         `json:"difficulty" gorm:"not null;check:difficulty IN ('easy', 'medium', 'hard')"`
	Status           string         `json:"status" gorm:"not null;default:'pending';check:status IN ('pending', 'in_progress', 'completed', 'approved', 'rejected')"`
	AssignedChildID  *uint          `json:"assigned_child_id,omitempty"`
	CreatedBy        string         `json:"created_by" gorm:"not null"`
	DueDate          *time.Time     `json:"due_date,omitempty"`
	CompletionProof  *string        `json:"completion_proof,omitempty" gorm:"type:text"`
	SubmittedAt      *time.Time     `json:"submitted_at,omitempty"`
	ApprovedAt       *time.Time     `json:"approved_at,omitempty"`
	RejectedAt       *time.Time     `json:"rejected_at,omitempty"`
	RejectionReason  *string        `json:"rejection_reason,omitempty"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"-" gorm:"index"`

	// 关联关系
	Creator       *User  `json:"creator,omitempty" gorm:"foreignKey:CreatedBy;references:WalletAddress"`
	AssignedChild *Child `json:"assigned_child,omitempty" gorm:"foreignKey:AssignedChildID;references:ID"`
}

func (Task) TableName() string {
	return "tasks"
}