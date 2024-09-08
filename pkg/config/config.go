package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Config struct {
	AppPort     string         `yaml:"app_port"`
	Database    DatabaseConfig `yaml:"database"`
	JWTSecret   string         `yaml:"jwt_secret"`
	Environment string         `yaml:"environment"`
}

func LoadConfig(path string) (*Config, error) {
	configFile, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать файл конфигурации: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, fmt.Errorf("не удалось распарсить YAML: %w", err)
	}

	return &config, nil
}
