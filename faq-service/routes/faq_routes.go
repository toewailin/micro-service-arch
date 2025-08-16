package routes

import (
	"faq-service/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// InitializeRoutes sets up the routes for the FAQ service
func InitializeRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/api/faqs", func(c *gin.Context) {
		handlers.GetFAQs(c, db)
	})
	r.POST("/api/faqs", func(c *gin.Context) {
		handlers.CreateFAQ(c, db)
	})
}
