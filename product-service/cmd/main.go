package main

import (
	"log"
	"product-service/models"
	"product-service/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Connect to database
	db, err := gorm.Open(sqlite.Open("product.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	// Migrate schema
	db.AutoMigrate(&models.Product{})

	// Initialize routes
	routes.InitializeRoutes(r, db)

	// Start the server
	if err := r.Run(":9091"); err != nil {
		log.Fatal("Product service failed to start:", err)
	}
}
