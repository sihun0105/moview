package models

import (
	"time"
)

type Workspace struct {
    ID        int           `gorm:"primaryKey"`
    Name      string
    URL       string
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `gorm:"index"`
    OwnerId   int
    
	Channels  []Channel `gorm:"foreignKey:WorkspaceId"`
    User      User      `gorm:"foreignKey:OwnerId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Workspace) TableName() string {
	return "Workspace"
}
