package dotenv

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labscool/mb-appointment-system/internal/environment"
	"github.com/labscool/mb-appointment-system/internal/platform/logger"
)

const (
	envExtension = ".env"
)

var (
	// Get current file full path from runtime
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	projectRootPath = filepath.Join(filepath.Dir(b), "../../../")
)

func LoadDotEnvFile() error {
	path := loadFilePath()
	err := godotenv.Load(path)

	if err != nil {
		logger.Errorf("error loading dotenv file: %s - error: %s", path, err.Error())
		return err
	}

	return nil
}

func loadFilePath() string {
	filepath := fmt.Sprintf("%s/%s%s", projectRootPath, environment.Get().String(), envExtension)
	return strings.ToLower(filepath)
}
