package agenda

import (
	"time"

	"github.com/labscool/mb-appointment-system/internal/domain"
)

type (
	AgendaRequest struct {
		UserID  uint           `json:"user_id" validate:"required"`
		Date    time.Time      `json:"date" validate:"required"`
		Details []agendaDetail `json:"details" validate:"required,dive"`
	}

	agendaDetail struct {
		Start    time.Time `json:"start" validate:"required"`
		End      time.Time `json:"end" validate:"required"`
		Services []uint    `json:"services" validate:"required"`
	}
)

func (a *AgendaRequest) ToDomain() *domain.Agenda {
	agendaDetails := make([]domain.AgendaDetail, 0)

	for _, ad := range a.Details {
		newDetail := domain.AgendaDetail{
			Start:    ad.Start,
			End:      ad.End,
			Services: ad.Services,
		}

		agendaDetails = append(agendaDetails, newDetail)
	}

	return &domain.Agenda{
		UserID:  a.UserID,
		Date:    a.Date,
		Details: agendaDetails,
	}
}
