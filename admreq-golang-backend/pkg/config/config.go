package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	*config
}

func GetConfig() (*Config, error) {
	config_path := "config/config.yaml"

	if _, err := os.Stat(config_path); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file does not exist: %s", config_path)
	}

	var conf config

	if err := cleanenv.ReadConfig(config_path, &conf); err != nil {
		return nil, fmt.Errorf("error while reading config file: %v", err)
	}

	return &Config{&conf}, nil
}
