package main

import (
	"github.com/gin-gonic/gin"
	"github.com/labscool/mb-appointment-system/cmd/api/app"
)

func main() {
	r := gin.Default()

	app.InitRoutes(r)

	r.Run()
}
