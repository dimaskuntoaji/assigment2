package model

type Item struct {
	ID          uint   `gorm:"primaryKey;column:item_id" json:"lineItemId" validate:"required" example:"1"`
	ItemCode    string `gorm:"not null;type:VARCHAR(50)" form:"itemCode" json:"itemCode" validate:"required" binding:"required" example:"ic-123"`
	Description string `gorm:"type:TEXT" form:"description" json:"description" binding:"required" example:"A random description"`
	Quantity    uint   `gorm:"not null" form:"quantity" json:"quantity" validate:"required" binding:"required" example:"1"`
	OrderID     uint   `gorm:"not null" json:"-" validate:"required"`
	Orders      Order  `gorm:"foreignKey:OrderID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}