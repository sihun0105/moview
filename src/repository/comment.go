package repository

import (
	"moview/src/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	CreateComment(comment *models.Comment) error
	GetCommentByID(id int) (*models.Comment, error)
	UpdateComment(comment *models.Comment) error
	DeleteComment(id int) error
}

type commentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{
		DB: db,
	}
}

func (r *commentRepository) CreateComment(comment *models.Comment) (error) {
	return r.DB.Create(comment).Error
}

func (r *commentRepository) GetCommentByID(id int) (*models.Comment, error) {
	var comment models.Comment
	if err := r.DB.First(&comment, id).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *commentRepository) UpdateComment(comment *models.Comment) (error) {
	return r.DB.Save(comment).Error
}

func (r *commentRepository) DeleteComment(id int) error {
	return r.DB.Delete(&models.Comment{}, id).Error
}