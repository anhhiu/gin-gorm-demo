package models

import (
	"gorm.io/gorm"
)

type Shoe struct {
	gorm.Model
	Name     string  `json:"name"`
	Size     int     `json:"size"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
