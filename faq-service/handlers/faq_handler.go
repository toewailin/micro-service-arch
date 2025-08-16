package handlers

import (
	"faq-service/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetFAQs handles GET requests to fetch all FAQs
func GetFAQs(c *gin.Context, db *gorm.DB) {
	var faqs []models.FAQ
	if err := db.Find(&faqs).Error; err != nil {
		c.JSON(500, gin.H{"message": "Error fetching FAQs"})
		return
	}
	c.JSON(200, faqs)
}

// CreateFAQ handles POST requests to create a new FAQ
func CreateFAQ(c *gin.Context, db *gorm.DB) {
	var faq models.FAQ
	if err := c.BindJSON(&faq); err != nil {
		c.JSON(400, gin.H{"message": "Invalid request"})
		return
	}
	if err := db.Create(&faq).Error; err != nil {
		c.JSON(500, gin.H{"message": "Error creating FAQ"})
		return
	}
	c.JSON(201, faq)
}
