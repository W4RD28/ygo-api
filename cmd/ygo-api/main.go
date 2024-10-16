package main

import (
	"log"

	"github.com/W4RD28/ygo-api/api/controllers"
	"github.com/W4RD28/ygo-api/internal/db"
	"github.com/W4RD28/ygo-api/internal/minio"
	"github.com/W4RD28/ygo-api/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func serveApp() {
	router := gin.Default()
	apiV1 := router.Group("/api/v1")
	{
		controllers.InitAuthRoutes(apiV1)
		controllers.InitCardRoutes(apiV1)
		controllers.InitImageRoutes(apiV1)
	}

	router.Run(":8080")
}

func main() {
	loadEnv()
	loadDatabase()
	loadMinio()
	serveApp()
}

func loadDatabase() {
	db.Connect()
	db.Database.AutoMigrate(&models.Card{}, &models.Image{}, &models.User{})
}

func loadMinio() {
	minio.Connect()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
