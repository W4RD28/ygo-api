package controllers

import (
	"github.com/W4RD28/ygo-api/api/handlers"
	"github.com/W4RD28/ygo-api/api/middlewares"
	"github.com/gin-gonic/gin"
)

func InitCardRoutes(router *gin.RouterGroup) {
	cardGroup := router.Group("/cards")
	{
		cardGroup.GET("/find", handlers.FindCards)
		cardGroup.GET("/find/:id", handlers.FindCardById)

		protected := cardGroup.Group("/")
		protected.Use(middlewares.JWTAuthMiddleware())
		{
			protected.POST("/add", handlers.AddCard)
			protected.PUT("/edit", handlers.EditCard)
			protected.DELETE("/delete", handlers.DeleteCard)
		}
	}
}
