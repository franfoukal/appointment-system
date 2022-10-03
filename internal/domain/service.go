package domain

import (
	"net/url"
	"time"

	"github.com/labscool/mb-appointment-system/cmd/api/presenter"
	"github.com/labscool/mb-appointment-system/db/models"
)

type (
	Service struct {
		ID          int64
		Name        string
		Duration    time.Duration
		Description string
		ImageURL    url.URL
	}
)

func (s *Service) ToPresenter() presenter.Service {
	return presenter.Service{
		ID:          s.ID,
		Name:        s.Name,
		Duration:    int64(s.Duration),
		Description: s.Description,
		ImageURL:    s.ImageURL.String(),
	}
}

func (s *Service) ToDBModel() *models.Service {
	return &models.Service{
		Name:        s.Name,
		Duration:    s.Duration,
		Description: s.Description,
		ImageURL:    s.ImageURL.String(),
	}
}

func ServiceFromDBModel(svc *models.Service) *Service {
	service := &Service{
		ID:          int64(svc.ID),
		Name:        svc.Name,
		Duration:    svc.Duration,
		Description: svc.Description,
	}

	if url, err := url.Parse(svc.ImageURL); err == nil {
		service.ImageURL = *url
	}

	return service
}
