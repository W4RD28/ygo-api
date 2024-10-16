package utils

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/W4RD28/ygo-api/internal/models"
	"github.com/W4RD28/ygo-api/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var privateKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(user models.User) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}

func ValidateJWT(c *gin.Context) error {
	token, err := getToken(c)
	if err != nil {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}

func CurrentUser(c *gin.Context) (*models.User, error) {
	token, err := getToken(c)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	id := claims["id"].(float64)
	user, err := services.FindUserById(uint(id))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func getToken(c *gin.Context) (*jwt.Token, error) {
	bearerToken := c.GetHeader("Authorization")
	if bearerToken == "" {
		return nil, errors.New("authorization header is required")
	}

	tokenString := strings.Split(bearerToken, " ")[1]
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})
}
