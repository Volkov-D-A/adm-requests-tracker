package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	models "github.com/volkov-d-a/adm-requests-tracker/internal/models"
)

type Config struct {
	*models.Config
}

func GetConfig() (*Config, error) {
	config_path := os.Getenv("CONFIG_PATH")
	if config_path == "" {
		return nil, fmt.Errorf("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(config_path); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file does not exist: %s", config_path)
	}

	var conf models.Config

	if err := cleanenv.ReadConfig(config_path, &conf); err != nil {
		return nil, fmt.Errorf("Error while reading config file: %v", err)
	}

	return &Config{&conf}, nil
}
