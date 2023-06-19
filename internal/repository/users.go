package repository

import (
	"errors"
	"fmt"

	"github.com/labscool/mb-appointment-system/db/models"
	"github.com/labscool/mb-appointment-system/internal/domain"
	customerror "github.com/labscool/mb-appointment-system/internal/feature/custom"
	"gorm.io/gorm"
)

type (
	UserRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	record := u.db.Create(&user)
	if record.Error != nil {
		return nil, fmt.Errorf("error saving user into db: %s", record.Error)
	}
	return user, nil
}

func (u *UserRepository) GetByEmail(email string) (*domain.User, error) {
	var user models.User
	record := u.db.Where("email = ?", email).First(&user)

	if errors.Is(record.Error, gorm.ErrRecordNotFound) {
		return nil, customerror.EntityNotFoundError("user not found")
	}

	if record.Error != nil {
		return nil, customerror.InternalServerAPIError("error getting user from database")
	}

	return user.ToDomain(), nil
}

func (u *UserRepository) GetByID(id uint) (*domain.User, error) {
	var user models.User
	record := u.db.First(&user, id)

	if errors.Is(record.Error, gorm.ErrRecordNotFound) {
		return nil, customerror.EntityNotFoundError(fmt.Sprintf("user with ID: %d not found", id))
	}

	if record.Error != nil {
		return nil, customerror.InternalServerAPIError("error getting user from database")
	}

	return user.ToDomain(), nil
}
