package models

import (
	"time"
)
type Comment struct {
    ID        int           `gorm:"primaryKey"`
    MovieId   int
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `gorm:"index"`
    UserNo    int
    Comment   string

    Movie Movie `gorm:"foreignKey:MovieId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
    User  User  `gorm:"foreignKey:UserNo;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Comment) TableName() string {
	return "Comment"
}