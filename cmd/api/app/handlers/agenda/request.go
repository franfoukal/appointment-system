package agenda

import (
	"strconv"
	"time"

	"github.com/labscool/mb-appointment-system/internal/domain"
)

type (
	AgendaRequest struct {
		UserID  uint           `json:"user_id" validate:"required"`
		Date    CustomTime     `json:"date" validate:"required"`
		Details []agendaDetail `json:"details" validate:"required,dive"`
	}

	agendaDetail struct {
		Start    CustomTime `json:"start" validate:"required"`
		End      CustomTime `json:"end" validate:"required"`
		Services []uint     `json:"services" validate:"required"`
	}

	CustomTime struct {
		time.Time
	}
)

func (a *AgendaRequest) ToDomain() *domain.Agenda {
	agendaDetails := make([]domain.AgendaDetail, 0)

	for _, ad := range a.Details {
		newDetail := domain.AgendaDetail{
			Start:    ad.Start.Time,
			End:      ad.End.Time,
			Services: ad.Services,
		}

		agendaDetails = append(agendaDetails, newDetail)
	}

	return &domain.Agenda{
		UserID:  a.UserID,
		Date:    a.Date.Time,
		Details: agendaDetails,
	}
}

func (c *CustomTime) UnmarshalJSON(input []byte) error {
	millis, err := strconv.ParseInt(string(input), 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(0, millis*int64(time.Millisecond))
	c.Time = tm
	return nil
}
