package controllers

import (
	pb "api-gateway/pb"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	userClient pb.UserServiceClient
}

func NewUserController(userClient pb.UserServiceClient) *UserController {
	return &UserController{
		userClient: userClient,
	}
}

func (ctrl *UserController) GetUser(c *gin.Context) {
	idStr := c.Param("id")

	// Convert idStr to uint32
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Create gRPC request
	req := &pb.GetUserRequest{
		Id: uint32(id),
	}

	// Call gRPC method
	res, err := ctrl.userClient.GetUser(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return response
	c.JSON(http.StatusOK, res.User)
}

func (ctrl *UserController) GetAllUsers(c *gin.Context) {
	req := &pb.GetAllUsersRequest{}
	res, err := ctrl.userClient.GetAllUsers(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res.Users)
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	var user pb.CreateUserRequest
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := ctrl.userClient.CreateUser(context.Background(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res.User)
}
