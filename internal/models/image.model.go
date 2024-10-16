package models

import "gorm.io/gorm"

// Image struct
type Image struct {
	gorm.Model
	URL    string `gorm:"size:255;not null" json:"url"`
	Name   string `gorm:"size:255;not null" json:"name"`
	CardID uint   `json:"card_id"`
}
