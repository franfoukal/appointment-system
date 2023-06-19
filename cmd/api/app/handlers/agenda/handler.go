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
		ctx := context.Background()
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

		agenda, err := a.manager.CreateAgenda(ctx, agendaReq.ToDomain())
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
