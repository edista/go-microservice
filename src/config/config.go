package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App      AppConfig      `yaml:"app"`
	Database DatabaseConfig `yaml:"database"`
}

type AppConfig struct {
	Name string `yaml:"name"`
}

type DatabaseConfig struct {
	Uri string `yaml:"uri"`
}

func NewConfig() (*Config, error) {
	config := &Config{}

	configPath := os.Getenv("APP_CONFIG_PATH")

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
