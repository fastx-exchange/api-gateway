package routes

import (
	"fastx-api/src/controllers/user"
	"github.com/gin-gonic/gin"
)

func InitializeUserRoutes(r *gin.Engine) {
	userController := user.NewUserController()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/user", userController.GetAllUsers)
		// Add more user-related routes here
	}
}
