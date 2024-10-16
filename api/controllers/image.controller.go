package controllers

import (
	"github.com/W4RD28/ygo-api/api/handlers"
	"github.com/W4RD28/ygo-api/api/middlewares"
	"github.com/gin-gonic/gin"
)

func InitImageRoutes(router *gin.RouterGroup) {
	imageGroup := router.Group("/images")
	{
		imageGroup.GET("/find", handlers.GetImagesHandler)
		imageGroup.GET("/find/:id", handlers.GetImageHandler)

		protected := imageGroup.Group("/")
		protected.Use(middlewares.JWTAuthMiddleware())
		{
			protected.POST("/upload", handlers.UploadImageHandler)
			protected.PUT("/edit", handlers.EditImageHandler)
			protected.DELETE("/delete", handlers.DeleteImageHandler)
		}
	}
}
