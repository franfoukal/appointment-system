package domain

import "golang.org/x/crypto/bcrypt"

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

func (user *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
