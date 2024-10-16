package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/W4RD28/ygo-api/internal/inputs"
	"github.com/W4RD28/ygo-api/internal/minio"
	"github.com/W4RD28/ygo-api/internal/models"
	"github.com/W4RD28/ygo-api/internal/services"
	"github.com/W4RD28/ygo-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func UploadImageHandler(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	name := c.Request.FormValue("name")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cardID, err := strconv.Atoi(c.Request.FormValue("card_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid card ID"})
		return
	}
	defer file.Close()

	bucketName := os.Getenv("MINIO_BUCKET")
	endpoint := os.Getenv("MINIO_ENDPOINT")
	objectName := utils.DeleteSpaces(name)
	contentType := header.Header.Get("Content-Type")

	bucketExists, err := minio.BucketExists(bucketName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check bucket"})
		return
	}
	if !bucketExists {
		err := minio.MakeBucket(bucketName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bucket"})
			return
		}
	}

	err = minio.UploadImage(bucketName, objectName, contentType, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
		return
	}

	image := &models.Image{
		Name:   utils.DeleteSpaces(objectName),
		URL:    fmt.Sprintf("%s/%s/%s", endpoint, bucketName, objectName),
		CardID: uint(cardID),
	}

	savedImage, err := services.SaveImage(image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": savedImage})
}

func GetImagesHandler(c *gin.Context) {
	query := c.Query("query")
	images, err := services.FindImages(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find images"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": images})
}

func GetImageHandler(c *gin.Context) {
	id := c.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image ID"})
		return
	}
	image, err := services.FindImageById(uint(uintID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": image})
}

func EditImageHandler(c *gin.Context) {
	var input inputs.ImageEditInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	image, err := services.FindImageById(input.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	image.Name = input.Name
	image.CardID = input.CardID

	updatedImage, err := services.UpdateImage(image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": updatedImage})
}

func DeleteImageHandler(c *gin.Context) {
	var input inputs.ImageDeleteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	image, err := services.FindImageById(input.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	err = minio.DeleteImage(os.Getenv("MINIO_BUCKET"), image.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete image"})
		return
	}

	err = services.DeleteImage(image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": image})
}
