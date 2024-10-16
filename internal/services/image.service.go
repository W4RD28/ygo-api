package services

import (
	"github.com/W4RD28/ygo-api/internal/db"
	"github.com/W4RD28/ygo-api/internal/models"
)

func SaveImage(image *models.Image) (*models.Image, error) {
	err := db.Database.Create(&image).Error
	if err != nil {
		return &models.Image{}, err
	}
	return image, err
}

func FindImageById(id uint) (*models.Image, error) {
	var image models.Image
	err := db.Database.First(&image, id).Error
	if err != nil {
		return &models.Image{}, err
	}
	return &image, err
}

func FindImages(query string) (*[]models.Image, error) {
	var images []models.Image
	err := db.Database.Where("name ILIKE ?", "%"+query+"%").Find(&images).Error
	if err != nil {
		return &[]models.Image{}, err
	}
	return &images, err
}

func EditImage(image *models.Image) (*models.Image, error) {
	err := db.Database.Save(&image).Error
	if err != nil {
		return &models.Image{}, err
	}
	return image, err
}

func DeleteImage(image *models.Image) error {
	err := db.Database.Delete(&image).Error
	if err != nil {
		return err
	}
	return err
}
