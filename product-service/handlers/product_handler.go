package handlers

import (
	"product-service/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetProducts handles GET requests to fetch products
func GetProducts(c *gin.Context, db *gorm.DB) {
	var products []models.Product
	if err := db.Find(&products).Error; err != nil {
		c.JSON(500, gin.H{"message": "Error fetching products"})
		return
	}
	c.JSON(200, products)
}

// CreateProduct handles POST requests to create a new product
func CreateProduct(c *gin.Context, db *gorm.DB) {
	var product models.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(400, gin.H{"message": "Invalid request"})
		return
	}
	if err := db.Create(&product).Error; err != nil {
		c.JSON(500, gin.H{"message": "Error creating product"})
		return
	}
	c.JSON(201, product)
}
