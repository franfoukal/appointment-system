package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine) {

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
