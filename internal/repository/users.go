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
	UserRepository struct{}
)

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	record := orm.Instance.Create(&user)
	if record.Error != nil {
		return nil, fmt.Errorf("error saving user into db: %s", record.Error)
	}
	return user, nil
}

func (u *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	record := orm.Instance.Where("email = ?", email).First(&user)

	if errors.Is(record.Error, gorm.ErrRecordNotFound) {
		return nil, customerror.EntityNotFoundError("user not found")
	}

	if record.Error != nil {
		return nil, customerror.InternalServerAPIError("error getting user from database")
	}

	return &user, nil
}

func (u *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	record := orm.Instance.First(&user, id)

	if errors.Is(record.Error, gorm.ErrRecordNotFound) {
		return nil, customerror.EntityNotFoundError(fmt.Sprintf("user with ID: %d not found", id))
	}

	if record.Error != nil {
		return nil, customerror.InternalServerAPIError("error getting user from database")
	}

	return &user, nil
}
