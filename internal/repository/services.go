package repository

import (
	"errors"
	"fmt"

	"github.com/labscool/mb-appointment-system/db/models"
	customerror "github.com/labscool/mb-appointment-system/internal/feature/custom"
	"github.com/labscool/mb-appointment-system/internal/platform/orm"
	"gorm.io/gorm"
)

type (
	ServiceRepository struct{}
)

func NewServiceRepository() *ServiceRepository {
	return &ServiceRepository{}
}

func (s *ServiceRepository) CreateService(service *models.Service) (*models.Service, error) {
	record := orm.Instance.Create(&service)
	if record.Error != nil {
		return nil, fmt.Errorf("error saving service into db: %s", record.Error)
	}

	return service, nil
}

func (s *ServiceRepository) GetServices() ([]*models.Service, error) {
	var services []*models.Service
	result := orm.Instance.Find(&services)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, customerror.EntityNotFoundError("servce not found")
	}

	if result.Error != nil {
		return nil, customerror.InternalServerAPIError("error getting service from database")
	}

	return services, nil
}

func (s *ServiceRepository) UpdateService(serviceID uint, serviceToUpdate *models.Service) (*models.Service, error) {
	service, err := s.GetServiceByID(serviceID)
	if err != nil {
		return nil, err
	}

	service.Name = serviceToUpdate.Name
	service.Duration = serviceToUpdate.Duration
	service.Description = serviceToUpdate.Description
	service.ImageURL = serviceToUpdate.ImageURL

	record := orm.Instance.Save(&service)
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
	record := orm.Instance.Delete(&service)
	if record.Error != nil {
		if errors.Is(record.Error, gorm.ErrRecordNotFound) {
			return customerror.EntityNotFoundError("service not found")
		}
		return fmt.Errorf("error deleting service from db: %s", record.Error)
	}

	return nil
}

func (s *ServiceRepository) GetServiceByID(serviceID uint) (*models.Service, error) {
	var service *models.Service
	result := orm.Instance.First(&service, serviceID)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, customerror.EntityNotFoundError("service not found")
	}

	if result.Error != nil {
		return nil, customerror.InternalServerAPIError("error getting services from database")
	}

	return service, nil
}
