package agenda

import (
	"context"
	"fmt"
	"time"

	"github.com/labscool/mb-appointment-system/db/models"
	"github.com/labscool/mb-appointment-system/internal/domain"
	"golang.org/x/exp/maps"
)

type (
	agendaRepository interface {
		CreateAgenda(agenda *domain.Agenda) (*domain.Agenda, error)
	}

	serviceRepository interface {
		MGetServiceByID(serviceIDs []int) ([]*models.Service, error)
	}

	userRepository interface {
		GetByID(id uint) (*models.User, error)
	}
	kvsClient interface {
		Set(key string, value interface{}, expiration_min int64) error
		MSet(keys []string, values []interface{}, expiration_min int64) error
	}

	AgendaFeature struct {
		agendaRepository  agendaRepository
		serviceRepository serviceRepository
		userRepository    userRepository
		kvsClient         kvsClient
	}
)

const (
	TIME_SLOT_TEMPLATE_KEY = "T%dE%d"
	TIME_SLOT_DURATION_MIN = 15
)

func NewAgendaFeature(agendaRepository agendaRepository, serviceRepository serviceRepository,
	userRepository userRepository, kvsClient kvsClient) *AgendaFeature {
	return &AgendaFeature{
		agendaRepository:  agendaRepository,
		serviceRepository: serviceRepository,
		userRepository:    userRepository,
		kvsClient:         kvsClient,
	}
}

func (a *AgendaFeature) CreateAgenda(ctx context.Context, agenda *domain.Agenda) (*domain.Agenda, error) {
	var ids []int
	for _, d := range agenda.Details {
		for _, s := range d.Services {
			ids = append(ids, int(s))
		}
	}

	_, err := a.serviceRepository.MGetServiceByID(ids)
	if err != nil {
		return nil, err
	}

	user, err := a.userRepository.GetByID(agenda.UserID)
	if err != nil {
		return nil, err
	}

	agendaSaved, err := a.agendaRepository.CreateAgenda(agenda)
	if err != nil {
		return nil, err
	}

	if err := a.createTimeslots(agenda, user.ToDomain()); err != nil {
		return nil, err
	}

	return agendaSaved, nil
}

func (a *AgendaFeature) createTimeslots(agenda *domain.Agenda, user *domain.User) error {

	timeSlots := make(map[string]interface{}, 0)

	for _, d := range agenda.Details {
		unixValues := generateSlotsUnixValues(d.Start, d.End)
		for _, uv := range unixValues {
			slotName := fmt.Sprintf(TIME_SLOT_TEMPLATE_KEY, uv, user.ID)
			timeSlots[slotName] = string(domain.TimeSlotsStatusType.FREE)
		}
	}

	if err := a.kvsClient.MSet(maps.Keys(timeSlots), maps.Values(timeSlots), 0); err != nil {
		return err
	}

	return nil
}

func generateSlotsUnixValues(start time.Time, end time.Time) []int64 {
	control := start
	slotNames := make([]int64, 0)

	for end.After(control) {
		slotNames = append(slotNames, control.Unix())
		control = control.Add(15 * time.Minute)
	}

	return slotNames
}
