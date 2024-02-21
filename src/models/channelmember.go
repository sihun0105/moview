package models

import (
	"time"
)

type ChannelMember struct {
    ChannelID int `gorm:"primaryKey"`
    UserId    int `gorm:"primaryKey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    
	User      User `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (ChannelMember) TableName() string {
	return "ChannelMember"
}