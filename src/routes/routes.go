package routes

import (
	pb "api-gateway/pb"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine, userClient pb.UserServiceClient) {
	InitializeUserRoutes(r, userClient)
}
