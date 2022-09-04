package app

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/labscool/mb-appointment-system/internal/environment"
	"github.com/labscool/mb-appointment-system/internal/feature/users"
	"github.com/labscool/mb-appointment-system/internal/platform/orm"
	"github.com/labscool/mb-appointment-system/internal/platform/sqlconnector"
	repository "github.com/labscool/mb-appointment-system/internal/repository/users"
)

type (
	Resources struct {
		Enforcer            casbin.Enforcer
		RegistrationFeature users.Registration
		AuthFeature         users.Auth
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
	case environment.EnvironmentType.Local:
		db, _ = sqlconnector.InitDBLocalConnection()
		orm.Connect(db)
		orm.Migrate()
		enforcer = configureEnforcer()
	}

	userRepository := repository.NewUserRepository()

	return &Resources{
		Enforcer:            *enforcer,
		RegistrationFeature: *users.NewUserRegistrationFeature(userRepository),
		AuthFeature:         *users.NewUserAuthFeature(userRepository),
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
