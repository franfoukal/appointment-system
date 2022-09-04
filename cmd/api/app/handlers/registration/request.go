package registration

import "github.com/labscool/mb-appointment-system/internal/domain"

type (
	UserRegistrationRequest struct {
		FirstName string `json:"first_name" binding:"required"`
		LastName  string `json:"last_name" binding:"required"`
		Username  string `json:"username" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required"`
		Role      string `json:"role" binding:"required"`
	}
)

func (u *UserRegistrationRequest) ToDomain() *domain.User {
	return &domain.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		Role:      u.Role,
	}
}
