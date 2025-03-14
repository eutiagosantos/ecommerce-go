package controller

import (
	"api-golang/internal/core/domain"
	"api-golang/internal/core/domain/auth"
	"api-golang/internal/core/usecases"
	"api-golang/internal/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	UserUseCases usecases.UserUseCases
}

func NewUserController(userUseCases *usecases.UserUseCases) *UserController {
	return &UserController{UserUseCases: *userUseCases}
}

func (u *UserController) Create(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := u.UserUseCases.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if err := u.UserUseCases.Create(&user); err == nil {
		c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
	}

}

func (u *UserController) Login(c *gin.Context) {
	var authInput auth.AuthInput

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	userFound, err := u.UserUseCases.FindByEmail(authInput.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(authInput.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
	}

	token, err := middlewares.JwtMiddleware()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
