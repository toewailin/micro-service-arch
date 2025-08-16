package routes

import (
	"user-service/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// InitializeRoutes sets up the routes for the user service
func InitializeRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/api/users", func(c *gin.Context) {
		handlers.GetUsers(c, db)
	})
	r.POST("/api/users", func(c *gin.Context) {
		handlers.CreateUser(c, db)
	})
}
