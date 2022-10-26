package presenter

import "time"

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
