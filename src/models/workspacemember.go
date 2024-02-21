package models

import (
	"time"
)

type WorkspaceMember struct {
    WorkspaceID int `gorm:"primaryKey"`
    UserId      int `gorm:"primaryKey"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
    LoggedInAt  *time.Time `gorm:"index"`
    
	User        User       `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (WorkspaceMember) TableName() string {
	return "WorkspaceMember"
}
