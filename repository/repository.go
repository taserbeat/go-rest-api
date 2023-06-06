package repository

import (
	"go-rest-api/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *models.User, email string) error
	CreateUser(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *models.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) CreateUser(user *models.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}