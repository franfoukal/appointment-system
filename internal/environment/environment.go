package environment

import (
	"os"
	"strings"
)

type Environment int

const (
	Prod Environment = iota
	Test
	Local
)

func (d Environment) String() string {
	return [...]string{"PRODUCTION", "TEST", "LOCAL"}[d]
}

func Get() Environment {
	switch strings.ToUpper(os.Getenv("ENVIRONMENT")) {
	case "PRODUCTION":
		return Prod
	case "TEST":
		return Test
	case "DEVELOPMENT":
		return Local
	default:
		return Local
	}
}
