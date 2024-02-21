package models

import (
	"time"
)

type Channel struct {
    ID          int       `gorm:"primaryKey"`
    Name        string
    Private     bool
    CreatedAt   time.Time
    UpdatedAt   time.Time
    WorkspaceId int
    
	Workspace   Workspace `gorm:"foreignKey:WorkspaceId"`
}

func (Channel) TableName() string {
    return "Channel"
}