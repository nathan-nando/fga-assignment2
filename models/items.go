package models

type Item struct {
	ID          string `gorm:"primaryKey" json:"item_id"`
	ItemCode    string `gorm:"type:varchar(190)" json:"item_code"`
	Description string `gorm:"type:varchar(190)" json:"description"`
	Quantity    int    `json:"quantity"`
	OrderId     string `json:"order_id" sql:"REFERENCES orders(id) ON DELETE CASCADE"`
}
