package models

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	Name      string `gorm:"size:255;not null;" json:"name"`
	Type      string `gorm:"size:255;not null;" json:"type"`
	Race      string `gorm:"size:255;not null;" json:"race"`
	Attribute string `gorm:"size:255;" json:"attribute"`

	Level       int    `gorm:"not null;" json:"level"`
	Attack      int    `json:"attack"`
	Defense     int    `json:"defense"`
	Description string `gorm:"size:255;not null;" json:"description"`

	Images []Image `gorm:"foreignKey:CardID" json:"images"`
}
