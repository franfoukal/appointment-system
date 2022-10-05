package services

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/labscool/mb-appointment-system/cmd/api/presenter"
	"github.com/labscool/mb-appointment-system/internal/domain"
	customerror "github.com/labscool/mb-appointment-system/internal/feature/custom"
	"github.com/labscool/mb-appointment-system/internal/platform/logger"
)

type (
	ServiceManager interface {
		CreateService(ctx context.Context, service *domain.Service) (*domain.Service, error)
		GetServices(ctx context.Context) ([]*domain.Service, error)
		UpdateService(ctx context.Context, serviceID uint, service *domain.Service) (*domain.Service, error)
		DeleteService(ctx context.Context, serviceID uint) error
	}
	ServiceHandler struct {
		manager ServiceManager
	}
)

func NewServiceHandler(manager ServiceManager) *ServiceHandler {
	return &ServiceHandler{manager: manager}
}

func (s *ServiceHandler) CreateService() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		var svcRequest ServiceRequest
		if err := c.ShouldBindJSON(&svcRequest); err != nil {
			logger.Errorf(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		if err := validator.New().Struct(&svcRequest); err != nil {
			logger.Errorf(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		svc, err := s.manager.CreateService(ctx, svcRequest.ToDomain())
		if err != nil {
			c.JSON(http.StatusInternalServerError, customerror.InternalServerAPIError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, svc.ToPresenter())
	}
}

func (s *ServiceHandler) GetServices() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		servicesDomain, err := s.manager.GetServices(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, customerror.InternalServerAPIError(err.Error()))
			return
		}

		services := make([]presenter.Service, 0)
		for _, ss := range servicesDomain {
			services = append(services, ss.ToPresenter())
		}

		c.JSON(http.StatusOK, services)
	}
}

func (s *ServiceHandler) UpdateService() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		var svcRequest ServiceRequest
		if err := c.ShouldBindJSON(&svcRequest); err != nil {
			logger.Errorf(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		if err := validator.New().Struct(&svcRequest); err != nil {
			logger.Errorf(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		serviceIDParam := c.Param("id")
		id, err := strconv.ParseInt(serviceIDParam, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, customerror.BadRequestAPIError(err.Error()))
			return
		}

		svc, err := s.manager.UpdateService(ctx, uint(id), svcRequest.ToDomain())
		if _, ok := err.(customerror.EntityNotFoundError); ok {
			c.JSON(http.StatusNotFound, customerror.NotFoundAPIError("There is not service with the id provided"))
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, customerror.InternalServerAPIError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, svc.ToPresenter())
	}
}

func (s *ServiceHandler) DeleteService() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		serviceIDParam := c.Param("id")
		id, err := strconv.ParseInt(serviceIDParam, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, customerror.BadRequestAPIError(err.Error()))
			return
		}

		err = s.manager.DeleteService(ctx, uint(id))
		if _, ok := err.(customerror.EntityNotFoundError); ok {
			c.JSON(http.StatusNotFound, customerror.NotFoundAPIError("There is not service with the id provided"))
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, customerror.InternalServerAPIError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "deleted successfully",
		})
	}
}