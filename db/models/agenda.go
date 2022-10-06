package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type (
	Agenda struct {
		gorm.Model
		UserID  uint
		Date    time.Time
		Details datatypes.JSON
	}
)
