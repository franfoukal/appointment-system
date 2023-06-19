package models

import (
	"net/url"
	"time"

	"github.com/labscool/mb-appointment-system/internal/domain"
	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	Name        string        `json:"name"`
	Duration    time.Duration `json:"duration"`
	Description string        `json:"description"`
	ImageURL    string        `json:"image_url"`
}

func ServiceModelFromDomain(service *domain.Service) *Service {
	return &Service{
		Name:        service.Name,
		Duration:    service.Duration,
		Description: service.Description,
		ImageURL:    service.ImageURL.String(),
	}
}

func (s *Service) ToDomain() *domain.Service {
	service := &domain.Service{
		ID:          int64(s.ID),
		Name:        s.Name,
		Duration:    s.Duration,
		Description: s.Description,
	}

	if url, err := url.Parse(s.ImageURL); err == nil {
		service.ImageURL = *url
	}

	return service
}
