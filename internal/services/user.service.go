package services

import (
	"html"
	"strings"

	"github.com/W4RD28/ygo-api/internal/db"
	"github.com/W4RD28/ygo-api/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SaveUser(user *models.User) (*models.User, error) {
	err := db.Database.Create(&user).Error
	if err != nil {
		return &models.User{}, err
	}
	return user, err
}

func BeforeSaveUser(user *models.User) (*gorm.DB, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil, err
}

func ComparePassword(user *models.User, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := db.Database.Where("username = ?", username).First(&user).Error
	if err != nil {
		return &models.User{}, err
	}
	return &user, nil
}

func FindUserById(id uint) (*models.User, error) {
	var user models.User
	err := db.Database.First(&user, id).Error
	if err != nil {
		return &models.User{}, err
	}
	return &user, nil
}
