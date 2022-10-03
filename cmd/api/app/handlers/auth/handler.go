package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labscool/mb-appointment-system/cmd/api/presenter"
	customerror "github.com/labscool/mb-appointment-system/internal/feature/custom"
	"github.com/labscool/mb-appointment-system/internal/feature/users"
)

type (
	AuthHandler struct {
		feature users.Auth
	}
)

func NewTokenHandler(feature users.Auth) *AuthHandler {
	return &AuthHandler{feature: feature}
}

func (a *AuthHandler) GenerateToken(c *gin.Context) {
	var request AuthRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, customerror.BadRequestAPIError("check fields in the body"))
		return
	}

	token, err := a.feature.Authenticate(request.Email, request.Password)
	if _, ok := err.(customerror.EntityNotFoundError); ok {
		c.JSON(http.StatusNotFound, customerror.NotFoundAPIError("There are no user with the provided email"))
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, customerror.InternalServerAPIError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, presenter.Token{JWT: token})
}
