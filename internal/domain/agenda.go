package domain

import (
	"time"
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
	TimeSlots struct {
		Slot   string         `json:"slot"`
		Status TimeSlotStatus `json:"status"`
	}

	TimeSlotStatus string

	timeSlotsStatusType struct {
		FREE  TimeSlotStatus
		TAKEN TimeSlotStatus
		TBC   TimeSlotStatus
	}
)

var (
	TimeSlotsStatusType = timeSlotsStatusType{
		FREE:  TimeSlotStatus("FREE"),
		TAKEN: TimeSlotStatus("TAKEN"),
		TBC:   TimeSlotStatus("TBC"),
	}
)
