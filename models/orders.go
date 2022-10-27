package models

import (
	"time"
)

type Order struct {
	ID           string    `gorm:"primaryKey" json:"order_id"`
	CustomerName string    `gorm:"type:varchar(190)" json:"customer_name" `
	Items        []Item    `json:"items"`
	OrderedAt    time.Time `json:"ordered_at"`
}
