package users

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labscool/mb-appointment-system/db/models"
	"github.com/labscool/mb-appointment-system/internal/domain"
	customerror "github.com/labscool/mb-appointment-system/internal/feature/custom"
)

type (
	JWTClaim struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		jwt.StandardClaims
	}

	Auth struct {
		userRepository
	}
)

var JWTKey = []byte(os.Getenv("JWT_KEY"))

func NewUserAuthFeature(repo userRepository) *Auth {
	return &Auth{userRepository: repo}
}

func (a *Auth) Authenticate(email, password string) (string, error) {
	user, err := a.checkCredentials(email, password)
	if err != nil {
		return "", err
	}

	payload := domain.TokenPayload{
		Username: user.Username,
		Email:    user.Email,
	}

	token, err := a.generateToken(payload)
	if err != nil {
		return "", customerror.InternalServerError("error generating token")
	}

	return token, nil
}

func (a *Auth) generateToken(payload domain.TokenPayload) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Username: payload.Username,
		Email:    payload.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(JWTKey)
}

func (a *Auth) ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(JWTKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}

func (a *Auth) checkCredentials(email, password string) (*models.User, error) {
	user, err := a.userRepository.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := user.CheckPassword(password); err != nil {
		return nil, customerror.ForbiddenError("invalid password")
	}

	return user, nil
}
