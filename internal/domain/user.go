package domain

import (
	"github.com/labscool/mb-appointment-system/cmd/api/presenter"
	"github.com/labscool/mb-appointment-system/db/models"
)

type (
	User struct {
		ID        uint
		FirstName string
		LastName  string
		Username  string
		Email     string
		Password  string
		Role      string
	}

	TokenPayload struct {
		ID       uint
		Username string
		Email    string
		Role     string
	}
)

func (u *User) ToDBModel() *models.User {
	return &models.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		Role:      u.Role,
	}
}

func UserFromDBModel(dbUser *models.User) *User {
	return &User{
		ID:        dbUser.ID,
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		Username:  dbUser.Username,
		Email:     dbUser.Email,
		Password:  dbUser.Password,
		Role:      dbUser.Role,
	}
}

func (u *User) ToPresenter() *presenter.Registration {
	return &presenter.Registration{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
		Role:      u.Role,
	}
}
