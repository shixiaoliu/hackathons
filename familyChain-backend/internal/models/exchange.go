package models

import "time"

// ExchangeStatus 表示兑换记录的状态
type ExchangeStatus string

const (
	// ExchangeStatusPending 表示待处理的兑换请求
	ExchangeStatusPending ExchangeStatus = "pending"
	// ExchangeStatusCompleted 表示已完成的兑换
	ExchangeStatusCompleted ExchangeStatus = "completed"
	// ExchangeStatusCancelled 表示已取消的兑换
	ExchangeStatusCancelled ExchangeStatus = "cancelled"
)

// Exchange 表示孩子兑换奖品的记录
type Exchange struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	RewardID      uint           `json:"reward_id" gorm:"not null;index"`
	ChildID       uint           `json:"child_id" gorm:"not null;index"`
	TokenAmount   int            `json:"token_amount" gorm:"not null"`
	Status        ExchangeStatus `json:"status" gorm:"type:varchar(20);default:'pending';index"`
	ExchangeDate  time.Time      `json:"exchange_date" gorm:"autoCreateTime"`
	CompletedDate *time.Time     `json:"completed_date"`
	Notes         string         `json:"notes" gorm:"type:text"`
	CreatedAt     time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"autoUpdateTime"`

	// 关联字段，不在数据库中
	RewardName  string `json:"reward_name,omitempty" gorm:"-"`
	RewardImage string `json:"reward_image,omitempty" gorm:"-"`
	ChildName   string `json:"child_name,omitempty" gorm:"-"`
}

// TableName 指定表名
func (Exchange) TableName() string {
	return "exchanges"
}

// ExchangeCreateRequest 表示创建兑换请求
type ExchangeCreateRequest struct {
	RewardID int    `json:"reward_id" binding:"required"`
	Notes    string `json:"notes"`
}

// ExchangeUpdateRequest 表示更新兑换状态的请求
type ExchangeUpdateRequest struct {
	Status ExchangeStatus `json:"status" binding:"required,oneof=pending completed cancelled"`
	Notes  string         `json:"notes"`
}
