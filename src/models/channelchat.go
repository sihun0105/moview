package models

import (
	"time"
)

type ChannelChat struct {
    ID        int       `gorm:"primaryKey"`
    Content   string
    CreatedAt time.Time
    UpdatedAt time.Time
    UserId    int
    ChannelID int
    
	User      User  `gorm:"foreignKey:UserId"`
    Channel   Channel `gorm:"foreignKey:ChannelID"`
}

func (ChannelChat) TableName() string {
	return "ChannelChat"
}