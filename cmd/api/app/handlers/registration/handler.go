package registration

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labscool/mb-appointment-system/internal/feature/users"
	"github.com/labscool/mb-appointment-system/internal/platform/logger"
)

type (
	RegistrationHandler struct {
		feature users.Registration
	}
)

func NewRegistrationHandler(feature users.Registration) *RegistrationHandler {
	return &RegistrationHandler{feature: feature}
}

func (r *RegistrationHandler) RegisterUser(c *gin.Context) {
	ctx := context.Background()
	var request UserRegistrationRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Errorf(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	newUser, err := r.feature.Register(ctx, request.ToDomain())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, newUser.ToPresenter())
}
