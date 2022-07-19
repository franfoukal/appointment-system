package main

import (
	"github.com/gin-gonic/gin"
	"github.com/labscool/mb-appointment-system/cmd/api/app"
	"github.com/labscool/mb-appointment-system/config"
	"github.com/labscool/mb-appointment-system/internal/platform/logger"
)

func main() {
	r := gin.Default()

	cfg, err := config.LoadConfiguration()
	if err != nil {
		panic(err)
	}

	logger.Infof("Configs: %+v", cfg)

	app.InitRoutes(r)

	r.Run()
}
