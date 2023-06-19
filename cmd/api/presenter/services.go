package presenter

import "github.com/labscool/mb-appointment-system/internal/domain"

type (
	Service struct {
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		Duration    int64  `json:"duration"`
		Description string `json:"description"`
		ImageURL    string `json:"image"`
	}
)

func ServiceFromDomain(service *domain.Service) Service {
	return Service{
		ID:          service.ID,
		Name:        service.Name,
		Duration:    int64(service.Duration),
		Description: service.Description,
		ImageURL:    service.ImageURL.String(),
	}
}
