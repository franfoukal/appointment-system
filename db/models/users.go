package models

import (
	"github.com/labscool/mb-appointment-system/internal/domain"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username" gorm:"unique"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	Role      string `json:"role" gorm:"-"`
}

type Role struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}

func UserModelFromDomain(user *domain.User) *User {
	return &User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		Role:      user.Role,
	}
}

func (u *User) ToDomain() *domain.User {
	return &domain.User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		Role:      u.Role,
	}
}
