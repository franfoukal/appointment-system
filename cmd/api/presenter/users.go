package presenter

import "github.com/labscool/mb-appointment-system/internal/domain"

type (
	Registration struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
		Email     string `json:"email"`
		Role      string `json:"role"`
	}

	Token struct {
		JWT string `json:"token"`
	}
)

func UserFromDomain(user *domain.User) *Registration {
	return &Registration{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
	}
}
