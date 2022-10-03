package models

import (
	"time"

	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	Name        string        `json:"name"`
	Duration    time.Duration `json:"duration"`
	Description string        `json:"description"`
	ImageURL    string        `json:"image_url"`
}
