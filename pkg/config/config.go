package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DB struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"db"`
	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
}

var (
	ConfigFilePath = "config.yml"
)

func LoadConfig() *Config {
	config := &Config{}
	data, err := os.ReadFile(ConfigFilePath)

	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	err = yaml.Unmarshal(data, config)

	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	return config
}
