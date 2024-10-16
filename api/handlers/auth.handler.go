package handlers

import (
	"net/http"

	input "github.com/W4RD28/ygo-api/internal/inputs"
	model "github.com/W4RD28/ygo-api/internal/models"
	"github.com/W4RD28/ygo-api/internal/services"
	"github.com/W4RD28/ygo-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input input.AuthRegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}
	services.BeforeSaveUser(&user)

	savedUser, err := services.SaveUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": savedUser})
}

func Login(c *gin.Context) {
	var input input.AuthLoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := services.FindUserByUsername(input.Username)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := services.ComparePassword(user, input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateJWT(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
