package users

import (
	"context"

	"github.com/labscool/mb-appointment-system/db/models"
	"github.com/labscool/mb-appointment-system/internal/domain"
	"github.com/labscool/mb-appointment-system/internal/platform/logger"
)

type (
	userRepository interface {
		CreateUser(user *models.User) (*models.User, error)
		GetByEmail(email string) (*models.User, error)
	}

	Registration struct {
		userRepository
	}
)

func NewUserRegistrationFeature(repo userRepository) *Registration {
	return &Registration{userRepository: repo}
}

func (r *Registration) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	userDBModel := models.UserModelFromDomain(user)
	if err := userDBModel.HashPassword(); err != nil {
		logger.Errorf("error hashing password: %s", err.Error())
		return nil, err
	}

	userDB, err := r.userRepository.CreateUser(userDBModel)
	if err != nil {
		logger.Errorf("error saving new user to DB: %s", err.Error())
		return nil, err
	}

	return userDB.ToDomain(), nil
}
