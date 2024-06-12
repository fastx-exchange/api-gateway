package user

import (
	"fastx-api/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	userService := services.NewUserService()
	return &UserController{userService: userService}
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
