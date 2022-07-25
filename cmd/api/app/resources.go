package app

import (
	"github.com/labscool/mb-appointment-system/internal/environment"
	"github.com/labscool/mb-appointment-system/internal/platform/logger"
	"github.com/labscool/mb-appointment-system/internal/platform/sqlconnector"
)

func BuildDependencies() {

	switch environment.Get() {
	case environment.EnvironmentType.Local:
		db, _ := sqlconnector.InitDBLocalConnection()
		logger.Infof("%+v", db)
	}
}
