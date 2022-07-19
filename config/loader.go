package config

import (
	"embed"
	"fmt"
	"strings"

	"github.com/labscool/mb-appointment-system/internal/environment"
	"github.com/labscool/mb-appointment-system/internal/platform/logger"
	"gopkg.in/yaml.v2"
)

const ymlExtension = ".yml"

//go:embed *.yml
var configFiles embed.FS

type Config struct {
	AppName string `yaml:"app-name"`
}

func LoadConfiguration() (Config, error) {
	filepath := loadFilePath()
	logger.Infof("Config loaded from: %s", filepath)
	content, err := configFiles.ReadFile(filepath)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read the file %v. %w", filepath, err)
	}

	return parse(content)
}

func loadFilePath() string {
	filepath := fmt.Sprintf("%s%s", environment.Get().String(), ymlExtension)
	return strings.ToLower(filepath)
}

func parse(content []byte) (Config, error) {
	var cfg Config

	if err := yaml.Unmarshal(content, &cfg); err != nil {
		return Config{}, fmt.Errorf("failed to parse the config. %w", err)
	}
	return cfg, nil
}
