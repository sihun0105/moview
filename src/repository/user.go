package repository

import (
	"moview/src/models"

	"gorm.io/gorm"
)

type UserRepository interface {
    CreateUser(user *models.User) error
    GetUserByID(id int) (*models.User, error)
    UpdateUser(user *models.User) error
    DeleteUser(id int) error
}
type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{
        DB: db,
    }
}

func (r *userRepository) CreateUser(user *models.User) error {
    return r.DB.Create(user).Error
}

func (r *userRepository) GetUserByID(id int) (*models.User, error) {
    var user models.User
    if err := r.DB.First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *userRepository) UpdateUser(user *models.User) error {
    return r.DB.Save(user).Error
}

func (r *userRepository) DeleteUser(id int) error {
    return r.DB.Delete(&models.User{}, id).Error
}
