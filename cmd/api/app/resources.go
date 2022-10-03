package app

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/labscool/mb-appointment-system/cmd/api/app/handlers/auth"
	"github.com/labscool/mb-appointment-system/cmd/api/app/handlers/registration"
	service_handler "github.com/labscool/mb-appointment-system/cmd/api/app/handlers/services"
	"github.com/labscool/mb-appointment-system/internal/environment"
	"github.com/labscool/mb-appointment-system/internal/feature/services"
	"github.com/labscool/mb-appointment-system/internal/feature/users"
	"github.com/labscool/mb-appointment-system/internal/platform/orm"
	"github.com/labscool/mb-appointment-system/internal/platform/sqlconnector"
	"github.com/labscool/mb-appointment-system/internal/repository"
)

type (
	Resources struct {
		Enforcer            casbin.Enforcer
		AuthHandler         auth.AuthHandler
		RegistrationHandler registration.RegistrationHandler
		ServiceHandler      service_handler.ServiceHandler
	}
)

var (
	// Get current file full path from runtime
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	projectRootPath = filepath.Join(filepath.Dir(b), "../../../")
)

func BuildDependencies() *Resources {

	var db *sql.DB
	var enforcer *casbin.Enforcer

	switch environment.Get() {
	case environment.Type.Local:
		db, _ = sqlconnector.InitDBLocalConnection()
		orm.Connect(db)
		orm.DevelopmentMigrations()
		enforcer = configureEnforcer()
	}

	// Repositories
	userRepository := repository.NewUserRepository()
	serviceRepository := repository.NewServiceRepository()

	// Features
	authFeature := users.NewUserAuthFeature(userRepository)
	registrationFeature := users.NewUserRegistrationFeature(userRepository)
	serviceFeature := services.NewServiceFeature(serviceRepository)

	return &Resources{
		Enforcer:            *enforcer,
		AuthHandler:         *auth.NewTokenHandler(*authFeature),
		RegistrationHandler: *registration.NewRegistrationHandler(*registrationFeature),
		ServiceHandler:      *service_handler.NewServiceHandler(serviceFeature),
	}
}

func configureEnforcer() *casbin.Enforcer {
	// Initialize  casbin adapter
	adapter, err := gormadapter.NewAdapterByDB(orm.Instance)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize casbin adapter: %v", err))
	}

	// Load model configuration file and policy store adapter
	enforcerPath := fmt.Sprintf("%s/%s", projectRootPath, "config/rbac/rbac_model.conf")
	enforcer, err := casbin.NewEnforcer(enforcerPath, adapter)

	if err != nil {
		panic(fmt.Sprintf("failed to create casbin enforcer: %v", err))
	}

	//add policy
	if hasPolicy := enforcer.HasPolicy("admin", "agenda", "create"); !hasPolicy {
		enforcer.AddPolicy("admin", "agenda", "create")
	}

	if hasPolicy := enforcer.HasPolicy("employee", "agenda", "create"); !hasPolicy {
		enforcer.AddPolicy("employee", "agenda", "create")
	}

	return enforcer
}
