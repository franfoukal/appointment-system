package agenda

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/labscool/mb-appointment-system/cmd/api/presenter"
	"github.com/labscool/mb-appointment-system/internal/domain"
	customerror "github.com/labscool/mb-appointment-system/internal/feature/custom"
	"github.com/labscool/mb-appointment-system/internal/platform/logger"
)

type (
	AgendaManager interface {
		CreateAgenda(ctx context.Context, agenda *domain.Agenda) (*domain.Agenda, error)
		GetAgendas(ctx context.Context) ([]*domain.Agenda, error)
	}
	AgendaHandler struct {
		manager AgendaManager
	}
)

func NewAgendaHandler(manager AgendaManager) *AgendaHandler {
	return &AgendaHandler{manager: manager}
}

func (a *AgendaHandler) CreateAgenda() gin.HandlerFunc {
	return func(c *gin.Context) {
		var agendaReq AgendaRequest
		if err := c.ShouldBindJSON(&agendaReq); err != nil {
			logger.Errorf(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if err := validator.New().Struct(&agendaReq); err != nil {
			logger.Errorf(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// TODO: delete timezone and hours part to dates before saving

		agenda, err := a.manager.CreateAgenda(c, agendaReq.ToDomain())
		if err != nil {
			logger.Errorf(err.Error())
			if _, ok := err.(customerror.EntityNotFoundError); ok {
				c.JSON(http.StatusNotFound, customerror.NotFoundAPIError(err.Error()))
				return
			}

			c.JSON(http.StatusInternalServerError, customerror.InternalServerAPIError(err.Error()))
			c.Abort()
			return
		}

		c.JSON(http.StatusCreated, presenter.AgendaFromDomain(agenda))
	}
}

func (a *AgendaHandler) GetAgendas() gin.HandlerFunc {
	return func(c *gin.Context) {
		results, err := a.manager.GetAgendas(c)
		if err != nil {
			logger.Errorf(err.Error())
			c.JSON(http.StatusInternalServerError, customerror.InternalServerAPIError(err.Error()))
			c.Abort()
			return
		}

		agendas := make([]*presenter.Agenda, 0)
		for _, agenda := range results {
			agendas = append(agendas, presenter.AgendaFromDomain(agenda))
		}

		c.JSON(http.StatusOK, agendas)
	}
}
