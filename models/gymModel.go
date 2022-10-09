package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Gym struct {
	gorm.Model

	Name        string          `json:"name" binding:"required" gorm:"type:text"`
	Images      pq.StringArray  `json:"images" gorm:"type:text[]"`
	Location    string          `json:"location" binding:"required" gorm:"type:text"`
	Coordinates pq.Float64Array `json:"coordinates" gorm:"type:float[]"`
	Description string          `json:"description" gorm:"type:text"`
	Logo        string          `json:"logo" gorm:"type:text"`
	Type        string          `json:"type" binding:"required" gorm:"type:text"`
	Amenities   pq.StringArray  `json:"amenities" binding:"required" gorm:"type:text[]"`
}
