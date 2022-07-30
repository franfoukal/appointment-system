package app

import (
	"github.com/labscool/mb-appointment-system/internal/environment"
	"github.com/labscool/mb-appointment-system/internal/feature/users"
	"github.com/labscool/mb-appointment-system/internal/platform/logger"
	"github.com/labscool/mb-appointment-system/internal/platform/orm"
	"github.com/labscool/mb-appointment-system/internal/platform/sqlconnector"
	repository "github.com/labscool/mb-appointment-system/internal/repository/users"
)

type (
	Resources struct {
		RegistrationFeature users.Registration
	}
)

func BuildDependencies() *Resources {

	switch environment.Get() {
	case environment.EnvironmentType.Local:
		db, _ := sqlconnector.InitDBLocalConnection()
		orm.Connect(db)
		orm.Migrate()
		logger.Infof("%+v", db)
	}

	userRepository := repository.NewUserRepository()

	return &Resources{
		RegistrationFeature: *users.NewUserRegistrationFeature(userRepository),
	}
}
