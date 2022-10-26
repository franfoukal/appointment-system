package domain

import (
	"encoding/json"
	"time"

	"github.com/labscool/mb-appointment-system/cmd/api/presenter"
	"github.com/labscool/mb-appointment-system/db/models"
	"gorm.io/datatypes"
)

type (
	Agenda struct {
		ID      uint
		UserID  uint
		Date    time.Time
		Details []AgendaDetail `json:"details"`
	}

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

func (a *Agenda) ToPresenter() *presenter.Agenda {
	details := make([]presenter.AgendaDetail, 0)
	for _, d := range a.Details {
		newDetail := presenter.AgendaDetail{
			Start:    d.Start,
			End:      d.End,
			Services: d.Services,
		}

		details = append(details, newDetail)
	}
	return &presenter.Agenda{
		ID:      a.ID,
		UserID:  a.UserID,
		Date:    a.Date,
		Details: details,
	}
}

func AgendaFromDBModel(model *models.Agenda) (*Agenda, error) {
	var details []AgendaDetail
	if err := json.Unmarshal(model.Details, &details); err != nil {
		return nil, err
	}

	return &Agenda{
		ID:      model.ID,
		UserID:  model.UserID,
		Date:    model.Date,
		Details: details,
	}, nil
}
