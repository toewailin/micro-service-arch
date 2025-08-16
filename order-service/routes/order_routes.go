package routes

import (
	"order-service/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// InitializeRoutes sets up the routes for the order service
func InitializeRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/api/orders", func(c *gin.Context) {
		handlers.GetOrders(c, db)
	})
	r.POST("/api/orders", func(c *gin.Context) {
		handlers.CreateOrder(c, db)
	})
	r.PUT("/api/orders/:id", func(c *gin.Context) {
		handlers.UpdateOrder(c, db)
	})
	r.DELETE("/api/orders/:id", func(c *gin.Context) {
		handlers.DeleteOrder(c, db)
	})
}
