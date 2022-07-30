package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labscool/mb-appointment-system/cmd/api/app/handlers/registration"
	"github.com/labscool/mb-appointment-system/cmd/api/presenter"
)

func InitRoutes(app *gin.Engine, resources *Resources) {

	registrationHandler := registration.NewRegistrationHandler(resources.RegistrationFeature)

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, presenter.PingResponse{Message: "pong"})
	})

	app.POST("/registration", registrationHandler.RegisterUser)
}
