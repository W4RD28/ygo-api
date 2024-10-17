package services

import (
	"github.com/W4RD28/ygo-api/internal/db"
	"github.com/W4RD28/ygo-api/internal/models"
)

func SaveCard(card *models.Card) (*models.Card, error) {
	err := db.Database.Create(&card).Error
	if err != nil {
		return &models.Card{}, err
	}
	return card, err
}

func EditCard(card *models.Card) (*models.Card, error) {
	err := db.Database.Save(&card).Error
	if err != nil {
		return &models.Card{}, err
	}
	return card, err
}

func DeleteCard(card *models.Card) error {
	err := db.Database.Delete(&card).Error
	if err != nil {
		return err
	}
	return err
}

func FindCardById(id uint) (*models.Card, error) {
	var card models.Card
	err := db.Database.Preload("Images").First(&card, id).Error
	if err != nil {
		return &models.Card{}, err
	}
	return &card, err
}

func FindCards(query string) (*[]models.Card, error) {
	var cards []models.Card

	err := db.Database.Preload("Images").Where("name LIKE ?", "%"+query+"%").
		Or("description LIKE ?", "%"+query+"%").Find(&cards).Error

	if err != nil {
		return &[]models.Card{}, err
	}
	return &cards, err
}
