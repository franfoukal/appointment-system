package repository

import (
	"fmt"

	"github.com/labscool/mb-appointment-system/db/models"
	"github.com/labscool/mb-appointment-system/internal/domain"
	customerror "github.com/labscool/mb-appointment-system/internal/feature/custom"
	"gorm.io/gorm"
)

type AgendaRepository struct {
	db *gorm.DB
}

func NewAgendaRepository(db *gorm.DB) *AgendaRepository {
	return &AgendaRepository{
		db: db,
	}
}

func (a *AgendaRepository) CreateAgenda(agenda *domain.Agenda) (*domain.Agenda, error) {
	model, err := models.AgendaModelFromDomain(agenda)
	if err != nil {
		errStr := fmt.Sprintf("error saving agenda into db: %s", err.Error())
		return nil, customerror.InternalServerError(errStr)
	}

	record := a.db.Create(&model)
	if record.Error != nil {
		return nil, customerror.InternalServerError(fmt.Sprintf("error saving agenda into db: %s", record.Error))
	}

	return model.ToDomain(), nil
}

func (a *AgendaRepository) GetAgendas() ([]*domain.Agenda, error) {
	agendas := make([]*models.Agenda, 0)
	result := a.db.Find(&agendas)

	if result.Error != nil {
		return nil, customerror.InternalServerAPIError("error getting service from database")
	}

	agendasDomain := make([]*domain.Agenda, 0)
	for _, agenda := range agendas {
		agendasDomain = append(agendasDomain, agenda.ToDomain())
	}

	return agendasDomain, nil
}
