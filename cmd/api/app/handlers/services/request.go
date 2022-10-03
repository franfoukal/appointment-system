package services

import (
	"net/url"
	"time"

	"github.com/labscool/mb-appointment-system/internal/domain"
)

type (
	ServiceRequest struct {
		ID          int64  `json:"id"`
		Name        string `json:"name" validate:"required,min=3,max=100"`
		Duration    int64  `json:"duration" validate:"required,gt=15,lt=540"` // minutes between 15' and 9h
		Description string `json:"description" validate:"omitempty,min=3,max=500"`
		ImageURL    string `json:"image" validate:"omitempty,url"`
	}
)

func (s *ServiceRequest) ToDomain() *domain.Service {
	service := &domain.Service{
		ID:          s.ID,
		Name:        s.Name,
		Duration:    time.Duration(s.Duration * 1000000000),
		Description: s.Description,
	}

	if url, err := url.Parse(s.ImageURL); err == nil {
		service.ImageURL = *url
	}

	return service
}
