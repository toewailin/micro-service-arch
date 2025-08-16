package routes

import (
	"product-service/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// InitializeRoutes sets up the routes for the product service
func InitializeRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/api/products", func(c *gin.Context) {
		handlers.GetProducts(c, db)
	})
	r.POST("/api/products", func(c *gin.Context) {
		handlers.CreateProduct(c, db)
	})
}
