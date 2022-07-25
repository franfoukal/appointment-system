package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labscool/mb-appointment-system/cmd/api/presenter"
)

func InitRoutes(app *gin.Engine) {

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, presenter.PingResponse{Message: "pong"})
	})
}
