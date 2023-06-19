package models

import (
	"encoding/json"
	"time"

	"github.com/labscool/mb-appointment-system/internal/domain"
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

func (a *Agenda) ToDomain() (*domain.Agenda, error) {
	var details []domain.AgendaDetail
	if err := json.Unmarshal(a.Details, &details); err != nil {
		return nil, err
	}

	return &domain.Agenda{
		ID:      a.ID,
		UserID:  a.UserID,
		Date:    a.Date,
		Details: details,
	}, nil
}

func AgendaModelFromDomain(agenda *domain.Agenda) (*Agenda, error) {
	detailsJSON, err := json.Marshal(agenda.Details)
	if err != nil {
		return nil, err
	}

	return &Agenda{
		UserID:  agenda.UserID,
		Date:    agenda.Date,
		Details: datatypes.JSON(detailsJSON),
	}, nil
}
