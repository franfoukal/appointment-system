package repository

import (
	"errors"
	"fmt"

	"github.com/labscool/mb-appointment-system/db/models"
	"github.com/labscool/mb-appointment-system/internal/domain"
	customerror "github.com/labscool/mb-appointment-system/internal/feature/custom"
	"gorm.io/gorm"
)

type (
	ServiceRepository struct {
		db *gorm.DB
	}
)

func NewServiceRepository(db *gorm.DB) *ServiceRepository {
	return &ServiceRepository{
		db: db,
	}
}

func (s *ServiceRepository) CreateService(service *domain.Service) (*domain.Service, error) {
	serviceDBEntity := models.ServiceModelFromDomain(service)
	record := s.db.Create(&serviceDBEntity)
	if record.Error != nil {
		return nil, fmt.Errorf("error saving service into db: %s", record.Error)
	}

	return serviceDBEntity.ToDomain(), nil
}

func (s *ServiceRepository) GetServices() ([]*domain.Service, error) {
	var servicesDB []*models.Service
	result := s.db.Find(&servicesDB)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, customerror.EntityNotFoundError("servce not found")
	}

	if result.Error != nil {
		return nil, customerror.InternalServerAPIError("error getting services from database")
	}

	services := make([]*domain.Service, 0)
	for _, service := range servicesDB {
		services = append(services, service.ToDomain())
	}

	return services, nil
}

func (s *ServiceRepository) UpdateService(serviceID uint, serviceToUpdate *domain.Service) (*domain.Service, error) {
	service, err := s.GetServiceByID(serviceID)
	if err != nil {
		return nil, err
	}

	service.Name = serviceToUpdate.Name
	service.Duration = serviceToUpdate.Duration
	service.Description = serviceToUpdate.Description
	service.ImageURL = serviceToUpdate.ImageURL

	record := s.db.Save(models.ServiceModelFromDomain(service))
	if errors.Is(record.Error, gorm.ErrRecordNotFound) {
		return nil, customerror.EntityNotFoundError("servce not found")
	}
	if record.Error != nil {
		return nil, fmt.Errorf("error updating service into db: %s", record.Error)
	}

	return service, nil
}

func (s *ServiceRepository) DeleteService(serviceID uint) error {
	service, err := s.GetServiceByID(serviceID)
	if err != nil {
		return err
	}
	record := s.db.Delete(models.ServiceModelFromDomain(service))
	if record.Error != nil {
		if errors.Is(record.Error, gorm.ErrRecordNotFound) {
			return customerror.EntityNotFoundError("service not found")
		}
		return fmt.Errorf("error deleting service from db: %s", record.Error)
	}

	return nil
}

func (s *ServiceRepository) GetServiceByID(serviceID uint) (*domain.Service, error) {
	var service *models.Service
	result := s.db.First(&service, serviceID)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, customerror.EntityNotFoundError("service not found")
	}

	if result.Error != nil {
		return nil, customerror.InternalServerAPIError("error getting services from database")
	}

	return service.ToDomain(), nil
}

func (s *ServiceRepository) MGetServiceByID(serviceIDs []int) ([]*domain.Service, error) {
	var foundServices []*models.Service
	result := s.db.Find(&foundServices, serviceIDs)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, customerror.EntityNotFoundError("service not found")
	}

	if result.Error != nil {
		return nil, customerror.InternalServerAPIError("error getting services from database")
	}

	services := make([]*domain.Service, 0)
	for _, service := range foundServices {
		services = append(services, service.ToDomain())
	}

	return services, nil
}
