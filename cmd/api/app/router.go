package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labscool/mb-appointment-system/cmd/api/app/middlewares"
	"github.com/labscool/mb-appointment-system/cmd/api/presenter"
)

func InitRoutes(app *gin.Engine, resources *Resources) {
	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, presenter.PingResponse{Message: "pong"})
	})

	app.POST("/registration", resources.RegistrationHandler.RegisterUser(&resources.Enforcer))
	app.POST("/login", resources.AuthHandler.GenerateToken)

	usersAPI := app.Group("/protected", middlewares.AuthJWT())
	usersAPI.GET("/example", middlewares.Authorize("agenda", "create", &resources.Enforcer), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})

	// CRUD services
	app.POST("/services", resources.ServiceHandler.CreateService())
	app.GET("/services", resources.ServiceHandler.GetServices())
	app.PUT("/services/:id", resources.ServiceHandler.UpdateService())
	app.DELETE("/services/:id", resources.ServiceHandler.DeleteService())

	//CRUD agenda
	app.POST("/agenda", resources.AgendaHandler.CreateAgenda())

}
