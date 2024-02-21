package models

import (
	"time"
)

type User struct {
    ID               int           `gorm:"primaryKey"`
    Email            string        `gorm:"unique;not null"`
    Nickname         string
    Password         string
    CreatedAt        time.Time
    UpdatedAt        time.Time
    DeletedAt        *time.Time `gorm:"index"`
    
    Comments         []Comment   `gorm:"foreignKey:UserNo"`
    ChannelChats     []ChannelChat `gorm:"foreignKey:UserId"`
    ChannelMembers   []ChannelMember `gorm:"foreignKey:UserId"`
    Workspaces       []Workspace `gorm:"foreignKey:OwnerId"`
    WorkspaceMembers []WorkspaceMember `gorm:"foreignKey:UserId"`
}

func (User) TableName() string {
    return "User"
}