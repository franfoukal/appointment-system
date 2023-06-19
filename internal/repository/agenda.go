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

	agendaResult, err := model.ToDomain()
	if err != nil {
		return nil, err
	}

	return agendaResult, nil
}
