package presenter

import (
	"time"

	"github.com/labscool/mb-appointment-system/internal/domain"
)

type (
	Agenda struct {
		ID      uint           `json:"id"`
		UserID  uint           `json:"user_id"`
		Date    time.Time      `json:"date"`
		Details []AgendaDetail `json:"details"`
	}

	AgendaDetail struct {
		Start    time.Time `json:"start"`
		End      time.Time `json:"end"`
		Services []uint    `json:"services"`
	}
)

func AgendaFromDomain(agenda *domain.Agenda) *Agenda {
	details := make([]AgendaDetail, 0)
	for _, d := range agenda.Details {
		newDetail := AgendaDetail{
			Start:    d.Start,
			End:      d.End,
			Services: d.Services,
		}

		details = append(details, newDetail)
	}
	return &Agenda{
		ID:      agenda.ID,
		UserID:  agenda.UserID,
		Date:    agenda.Date,
		Details: details,
	}
}
