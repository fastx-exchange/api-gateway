package routes

import (
	"fastx-api/src/controllers/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeUserRoutes(r *gin.Engine, db *gorm.DB) {
	userController := user.NewUserController(db)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/user", userController.GetAllUsers)
		// Add more user-related routes here
	}
}
