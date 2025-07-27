package controller

import (
	"donasi-yuk/helper"
	"donasi-yuk/model"
	"donasi-yuk/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService user.Service
}

func NewUserController(userService user.Service) *UserController {
	return &UserController{userService}
}

// REGISTER USER handler
func (uc *UserController) Register(c *gin.Context) {
	var input model.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	newUser, err := uc.userService.Register(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate JWT token
	token, err := helper.GenerateToken(int(newUser.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// ✅ Kirim response JSON yang benar
	c.JSON(http.StatusOK, gin.H{
		"user":  newUser,
		"token": token,
	})
}

// LOGIN USER handler
func (uc *UserController) Login(c *gin.Context) {
	var input model.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input!!"})
		return
	}

	loggedInUser, err := uc.userService.Login(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials!!"})
		return
	}

	token, err := helper.GenerateToken(int(loggedInUser.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed generate token!!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": loggedInUser, "token": token})
}
