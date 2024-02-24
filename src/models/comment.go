package models

import (
	"time"
)
type Comment struct {
    ID        int   `json:"cateno" gorm:"primaryKey"`
    MovieId   int   `json:"movieid" gorm:"column:movieId;not null"`
    CreatedAt time.Time     `json:"created_at" gorm:"column:createdAt;type:datetime;default:CURRENT_TIMESTAMP(6)"`
    UpdatedAt time.Time     `json:"updated_at" gorm:"column:updatedAt;type:datetime;default:CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6)"`
    DeletedAt *time.Time    `json:"deleted_at" gorm:"column:deletedAt;type:datetime;default:NULL;"`
    UserNo    int  `json:"userno" gorm:"column:userno;not null"`
    Comment   string `json:"comment" gorm:"column:comment;not null;"`

    Movie Movie `gorm:"foreignKey:MovieId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
    User  User  `gorm:"foreignKey:UserNo;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Comment) TableName() string {
	return "Comment"
}