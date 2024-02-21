package models

import (
	"time"
)

type Movie struct {
    ID        int           `gorm:"primaryKey"`
    Audience  int
    MovieCd   int           `gorm:"unique"`
    Title     string
    Rank      int
    CreatedAt time.Time
    UpdatedAt time.Time
    
	Comments  []Comment `gorm:"foreignKey:MovieId"`
}

func (Movie) TableName() string {
	return "Movie"
}
