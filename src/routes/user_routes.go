package routes

import (
	"api-gateway/pb"
	"api-gateway/src/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeUserRoutes(r *gin.Engine, userClient pb.UserServiceClient) {
	userController := controllers.NewUserController(userClient)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", userController.GetAllUsers)
		v1.GET("/users/:id", userController.GetUser)
		v1.POST("/users", userController.CreateUser)
	}
}
