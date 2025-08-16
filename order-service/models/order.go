package models

// Order represents the structure of an order entity
type Order struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
	Status    string  `json:"status"` // e.g., "pending", "shipped", "delivered"
}
