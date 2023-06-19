package users

import (
	"context"

	"github.com/labscool/mb-appointment-system/internal/domain"
	"github.com/labscool/mb-appointment-system/internal/platform/logger"
)

type (
	userRepository interface {
		CreateUser(user *domain.User) (*domain.User, error)
		GetByEmail(email string) (*domain.User, error)
	}

	Registration struct {
		userRepository
	}
)

func NewUserRegistrationFeature(repo userRepository) *Registration {
	return &Registration{userRepository: repo}
}

func (r *Registration) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	if err := user.HashPassword(); err != nil {
		logger.Errorf("error hashing password: %s", err.Error())
		return nil, err
	}

	createdUser, err := r.userRepository.CreateUser(user)
	if err != nil {
		logger.Errorf("error saving new user to DB: %s", err.Error())
		return nil, err
	}

	return createdUser, nil
}
