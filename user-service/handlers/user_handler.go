package handlers

import (
	"user-service/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetUsers handles GET requests to fetch users
func GetUsers(c *gin.Context, db *gorm.DB) {
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		c.JSON(500, gin.H{"message": "Error fetching users"})
		return
	}
	c.JSON(200, users)
}

// CreateUser handles POST requests to create a new user
func CreateUser(c *gin.Context, db *gorm.DB) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message": "Invalid request"})
		return
	}
	if err := db.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"message": "Error creating user"})
		return
	}
	c.JSON(201, user)
}
