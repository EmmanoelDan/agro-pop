package repositories

import (
	"github.com/EmmanoelDan/agro-pop/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (repo *UserRepository) Create(user *models.User) error {
	return repo.DB.Create(user).Error
}

func (repo *UserRepository) FindByUsername(username string) (*models.User,error) {
	var user models.User

	err := repo.DB.Where("username = ?", username).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}