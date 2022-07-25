package environment

import (
	"os"
	"strings"
)

type (
	Environment string

	environmentType struct {
		Test       Environment
		Local      Environment
		Production Environment
	}
)

var (
	EnvironmentType = environmentType{
		Test:       Environment("test"),
		Local:      Environment("local"),
		Production: Environment("production"),
	}
)

func (env Environment) String() string {
	return string(env)
}

func Get() Environment {
	env := os.Getenv("ENVIRONMENT")
	switch strings.ToLower(env) {
	case EnvironmentType.Local.String():
		return EnvironmentType.Local
	case EnvironmentType.Test.String():
		return EnvironmentType.Test
	case EnvironmentType.Production.String():
		return EnvironmentType.Production
	default:
		return EnvironmentType.Local
	}
}
