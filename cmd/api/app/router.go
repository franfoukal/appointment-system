package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labscool/mb-appointment-system/cmd/api/app/handlers/auth"
	"github.com/labscool/mb-appointment-system/cmd/api/app/handlers/registration"
	"github.com/labscool/mb-appointment-system/cmd/api/app/middlewares"
	"github.com/labscool/mb-appointment-system/cmd/api/presenter"
)

func InitRoutes(app *gin.Engine, resources *Resources) {

	registrationHandler := registration.NewRegistrationHandler(resources.RegistrationFeature)
	authHandler := auth.NewTokenHandler(resources.AuthFeature)

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, presenter.PingResponse{Message: "pong"})
	})

	app.POST("/registration", registrationHandler.RegisterUser(&resources.Enforcer))
	app.POST("/login", authHandler.GenerateToken)

	usersAPI := app.Group("/protected", middlewares.AuthJWT())
	usersAPI.GET("/example", middlewares.Authorize("agenda", "create", &resources.Enforcer), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})
}
