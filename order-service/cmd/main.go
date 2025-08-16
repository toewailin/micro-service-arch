package main

import (
	"log"
	"order-service/config"
	"order-service/models"
	"order-service/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Load configuration
	cfg := config.LoadConfig()

	// Connect to SQLite database
	db, err := gorm.Open(sqlite.Open(cfg.DatabaseFile), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Migrate schema
	err = db.AutoMigrate(&models.Order{})
	if err != nil {
		log.Fatal("Error migrating database:", err)
	}

	// Set up routes
	routes.InitializeRoutes(r, db)

	// Start the server
	if err := r.Run(":9092"); err != nil {
		log.Fatal("Order service failed to start:", err)
	}
}
