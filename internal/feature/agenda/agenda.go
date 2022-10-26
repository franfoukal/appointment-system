package agenda

import (
	"context"

	"github.com/labscool/mb-appointment-system/db/models"
	"github.com/labscool/mb-appointment-system/internal/domain"
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

	AgendaFeature struct {
		agendaRepository  agendaRepository
		serviceRepository serviceRepository
		userRepository    userRepository
	}
)

func NewAgendaFeature(agendaRepository agendaRepository,
	serviceRepository serviceRepository,
	userRepository userRepository) *AgendaFeature {
	return &AgendaFeature{
		agendaRepository:  agendaRepository,
		serviceRepository: serviceRepository,
		userRepository:    userRepository,
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

	_, err = a.userRepository.GetByID(agenda.UserID)
	if err != nil {
		return nil, err
	}

	agendaSaved, err := a.agendaRepository.CreateAgenda(agenda)
	if err != nil {
		return nil, err
	}

	return agendaSaved, nil
}
