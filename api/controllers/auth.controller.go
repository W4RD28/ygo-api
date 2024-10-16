package controllers

import (
	"github.com/W4RD28/ygo-api/api/handlers"
	"github.com/gin-gonic/gin"
)

func InitAuthRoutes(router *gin.RouterGroup) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", handlers.Register)
		authGroup.POST("/login", handlers.Login)
	}
}
