package entity

import (
	"time"
)

type Item struct {
	ItemID      uint `gorm:"primary_key"`
	ItemCode    string
	Description string
	Quantity    int
	OrderID     uint
	CustomGormModel
}

type CustomGormModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type ItemRqResponse struct {
	ItemID      uint   `json:"itemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
