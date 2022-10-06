package domain

import (
	"encoding/json"
	"time"

	"github.com/labscool/mb-appointment-system/db/models"
	"gorm.io/datatypes"
)

type (
	Agenda struct {
		ID      uint
		UserID  uint
		Date    time.Time
		Details AgendaDetails `json:"details"`
	}

	AgendaDetails []AgendaDetail

	AgendaDetail struct {
		Start    time.Time `json:"start"`
		End      time.Time `json:"end"`
		Services []uint    `json:"services"`
	}
)

func (a *Agenda) ToDBModel() (*models.Agenda, error) {
	detailsJSON, err := json.Marshal(a.Details)
	if err != nil {
		return nil, err
	}

	return &models.Agenda{
		UserID:  a.UserID,
		Date:    a.Date,
		Details: datatypes.JSON(detailsJSON),
	}, nil
}
