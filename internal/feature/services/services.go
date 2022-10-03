package services

import (
	"context"

	"github.com/labscool/mb-appointment-system/db/models"
	"github.com/labscool/mb-appointment-system/internal/domain"
	"github.com/labscool/mb-appointment-system/internal/platform/logger"
)

type (
	serviceRepository interface {
		CreateService(service *models.Service) (*models.Service, error)
		GetServices() ([]*models.Service, error)
		GetServiceByID(serviceID uint) (*models.Service, error)
		UpdateService(serviceID uint, service *models.Service) (*models.Service, error)
		DeleteService(serviceID uint) error
	}

	ServiceFeature struct {
		repository serviceRepository
	}
)

func NewServiceFeature(repository serviceRepository) *ServiceFeature {
	return &ServiceFeature{
		repository: repository,
	}
}

func (s *ServiceFeature) CreateService(ctx context.Context, newService *domain.Service) (*domain.Service, error) {
	service, err := s.repository.CreateService(newService.ToDBModel())
	if err != nil {
		logger.Errorf("error saving new service into DB: %s", err.Error())
		return nil, err
	}

	return domain.ServiceFromDBModel(service), nil
}

func (s *ServiceFeature) GetServices(ctx context.Context) ([]*domain.Service, error) {
	serviceList, err := s.repository.GetServices()
	if err != nil {
		logger.Errorf("error retrieving services from DB: %s", err.Error())
		return nil, err
	}

	var services []*domain.Service
	for _, ss := range serviceList {
		services = append(services, domain.ServiceFromDBModel(ss))
	}

	return services, nil
}

func (s *ServiceFeature) UpdateService(ctx context.Context, serviceID uint, serviceToUpdate *domain.Service) (*domain.Service, error) {
	service, err := s.repository.UpdateService(serviceID, serviceToUpdate.ToDBModel())
	if err != nil {
		logger.Errorf("error updating service into DB: %s", err.Error())
		return nil, err
	}
	return domain.ServiceFromDBModel(service), nil
}

func (s *ServiceFeature) DeleteService(ctx context.Context, serviceID uint) error {
	err := s.repository.DeleteService(serviceID)
	if err != nil {
		return err
	}
	return nil
}
