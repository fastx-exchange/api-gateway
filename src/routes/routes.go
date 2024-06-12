package routes

import (
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	InitializeUserRoutes(r)
}
