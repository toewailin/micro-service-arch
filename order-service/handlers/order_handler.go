package handlers

import (
	"order-service/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetOrders handles GET requests to fetch all orders
func GetOrders(c *gin.Context, db *gorm.DB) {
	var orders []models.Order
	if err := db.Find(&orders).Error; err != nil {
		c.JSON(500, gin.H{"message": "Error fetching orders"})
		return
	}
	c.JSON(200, orders)
}

// CreateOrder handles POST requests to create a new order
func CreateOrder(c *gin.Context, db *gorm.DB) {
	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(400, gin.H{"message": "Invalid request"})
		return
	}
	// Calculate total price (you could do this more intelligently by querying a product database)
	order.TotalPrice = float64(order.Quantity) * 10.0 // Assume each product costs 10 for simplicity
	if err := db.Create(&order).Error; err != nil {
		c.JSON(500, gin.H{"message": "Error creating order"})
		return
	}
	c.JSON(201, order)
}

// UpdateOrder handles PUT requests to update an existing order
func UpdateOrder(c *gin.Context, db *gorm.DB) {
	var order models.Order
	id := c.Param("id")
	if err := db.First(&order, id).Error; err != nil {
		c.JSON(404, gin.H{"message": "Order not found"})
		return
	}

	if err := c.BindJSON(&order); err != nil {
		c.JSON(400, gin.H{"message": "Invalid request"})
		return
	}

	// Update order status
	db.Save(&order)
	c.JSON(200, order)
}

// DeleteOrder handles DELETE requests to delete an order
func DeleteOrder(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	if err := db.Delete(&models.Order{}, id).Error; err != nil {
		c.JSON(500, gin.H{"message": "Error deleting order"})
		return
	}
	c.JSON(204, nil) // No content
}
