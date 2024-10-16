package models

import "gorm.io/gorm"

// Image struct
type Image struct {
	gorm.Model
	URL    string `gorm:"size:255;not null" json:"url"`
	CardID uint   `gorm:"not null" json:"card_id"`
}
