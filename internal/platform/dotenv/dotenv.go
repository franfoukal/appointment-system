package dotenv

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labscool/mb-appointment-system/internal/environment"
)

const envExtension = ".env"

func LoadDotEnvFile() error {
	return godotenv.Load(loadFilePath())
}

func loadFilePath() string {
	filepath := fmt.Sprintf("%s%s", environment.Get().String(), envExtension)
	return strings.ToLower(filepath)
}
