package models

import (
	"time"
)

type Movie struct {
    ID        int           `json:"MovieId" gorm:"primaryKey"`
    Audience  int           `json:"Audience" gorm:"column:audience;not null"`
    MovieCd   int           `json:"movieCd" gorm:"column:movieCd;not null;unique"`
    Title     string        `json:"title" gorm:"column:title;not null"`
    Rank      int           `json:"rank" gorm:"column:rank;not null"`
    CreatedAt time.Time     `json:"created_at" gorm:"column:createdAt;type:datetime;default:CURRENT_TIMESTAMP(6)"`
    UpdatedAt time.Time     `json:"updated_at" gorm:"column:updatedAt;type:datetime;default:CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6)"`
    
	Comments  []Comment `gorm:"foreignKey:MovieId"`
}

func (Movie) TableName() string {
	return "Movie"
}
