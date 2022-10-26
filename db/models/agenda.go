package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type (
	Agenda struct {
		gorm.Model
		UserID  uint      `gorm:"index:idx_user_date"`
		Date    time.Time `gorm:"index:idx_user_date"`
		Details datatypes.JSON
	}
)
