package model

import "time"

// Model represents the model for an order
type Order struct {
	ID           uint      `gorm:"primaryKey;column:order_id" json:"orderId" validate:"required"`
	CustomerName string    `gorm:"type:VARCHAR(50)" form:"customerName" json:"customerName" validate:"required" binding:"required"`
	OrderedAt    time.Time `gorm:"not null;type:timestamp;autoCreateTime" json:"orderedAt" validate:"required"`
	Items        []Item    `form:"items" json:"items" validate:"required" binding:"required"`
}