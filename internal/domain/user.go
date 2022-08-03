package domain

import (
	"github.com/labscool/mb-appointment-system/cmd/api/presenter"
	"github.com/labscool/mb-appointment-system/db/models"
)

type (
	User struct {
		FirstName string
		LastName  string
		Username  string
		Email     string
		Password  string
	}

	TokenPayload struct {
		Username string
		Email    string
	}
)

func (u *User) ToDBModel() *models.User {
	return &models.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
	}
}

func (u *User) ToPresenter() *presenter.Registration {
	return &presenter.Registration{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
	}
}
