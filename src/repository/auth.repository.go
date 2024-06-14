package repository

import (
	"moview/src/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	LogIn(user *models.User) (*models.User, error)
}

type authRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{
		DB: db,
	}
}

func (r *authRepository) LogIn(user *models.User) (*models.User, error) {
	var u models.User
	if err := r.DB.Where("email = ?", user.Email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}