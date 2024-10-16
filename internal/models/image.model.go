package models

import "gorm.io/gorm"

// Image struct
type Image struct {
	gorm.Model
	URL       string `gorm:"size:255;not null" json:"url"`
	Name      string `gorm:"size:255;not null" json:"name"`
	ArtNumber string `gorm:"size:255;not null" json:"art_number"`
	CardID    uint   `gorm:"not null" json:"card_id"`
}
