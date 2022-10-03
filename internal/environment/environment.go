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
	Type = environmentType{
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
	case Type.Local.String():
		return Type.Local
	case Type.Test.String():
		return Type.Test
	case Type.Production.String():
		return Type.Production
	default:
		return Type.Local
	}
}
