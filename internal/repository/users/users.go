package repository

import (
	"fmt"

	"github.com/labscool/mb-appointment-system/db/models"
	"github.com/labscool/mb-appointment-system/internal/platform/orm"
)

type (
	UserRepository struct{}
)

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) CreateUser(user *models.User) error {
	record := orm.Instance.Create(user)
	if record.Error != nil {
		return fmt.Errorf("error saving user into db: %s", record.Error)
	}
	return nil
}
