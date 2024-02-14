package models

import (
	"time"
)

type User struct {
    ID        int       `gorm:"primaryKey" json:"id"`
    Email     string    `json:"email"`
    Nickname  string    `json:"nickname"`
    Password  string    `json:"password"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
    DeletedAt time.Time `gorm:"index" json:"deletedAt"`
}

func (User) TableName() string {
    return "User"
}