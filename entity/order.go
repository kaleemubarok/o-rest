package entity

import (
	"time"
)

type Order struct {
	OrderID      uint `gorm:"primary_key"`
	CustomerName string
	OrderAt      time.Time
	CustomGormModel
}

type OrderRqResponse struct {
	OrderID      uint             `json:"orderId"`
	OrderAt      time.Time        `json:"orderAt" binding:"required"`
	CustomerName string           `json:"customerName" binding:"required"`
	Items        []ItemRqResponse `json:"items"`
}
