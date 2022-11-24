package main

import (
	"github.com/gin-gonic/gin"
	"github.com/labscool/mb-appointment-system/cmd/api/app"
	"github.com/labscool/mb-appointment-system/config"
	"github.com/labscool/mb-appointment-system/internal/environment"
	"github.com/labscool/mb-appointment-system/internal/platform/dotenv"
	"github.com/labscool/mb-appointment-system/internal/platform/logger"
)

func main() {
	r := gin.Default()

	if err := dotenv.LoadDotEnvFile(); err != nil {
		panic(err)
	}

	cfg, err := config.LoadConfiguration()
	if err != nil {
		panic(err)
	}

	resources := app.BuildDependencies(cfg)

	// DEBUG
	env := environment.Get()
	logger.Infof("env: %s", env)
	logger.Infof("Configs: %+v", cfg)

	app.InitRoutes(r, resources)

	r.Run()
}
