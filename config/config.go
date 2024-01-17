package config

import (
	"errors"
	"github.com/jinzhu/configor"
)

const (
	DEV  = "development"
	STG  = "staging"
	PROD = "production"
)

type Config struct {
	AppName string `default:"monorepo" env:"APP_NAME" required:"true" yaml:"app_name"`
	Env     string `default:"development" env:"ENV" yaml:"env"`

	Database Database `yaml:"database"`
	Server   Server   `yaml:"server"`
}

type DoctorConfig struct {
	AppName string `default:"monorepo" env:"APP_NAME" required:"true" yaml:"app_name"`

	Database Database `yaml:"database"`
}

// Database - Database configuration for Postgres
type Database struct {
	Host     string `default:"localhost" env:"DB_HOST" yaml:"host"`
	Port     string `default:"5432" env:"DB_PORT" yaml:"port"`
	User     string `default:"postgres" env:"DB_USER" yaml:"user"`
	Password string `default:"postgres" env:"DB_PASSWORD" yaml:"password"`
	Name     string `default:"monorepo" env:"DB_NAME" yaml:"name"`
}

// Server - Server configuration for HTTP
type Server struct {
	Port            string `default:"8080" env:"PORT" yaml:"port"`
	ShutdownTimeout int    `default:"10" env:"SHUTDOWN_TIMEOUT" yaml:"shutdown_timeout"`
}

func LoadConfig() (*Config, error) {
	var cfg Config

	err := configor.Load(&cfg, "config.yml")
	if err != nil {
		return nil, errors.New("ERROR: Failed to load config file %w" + err.Error())
	}

	return &cfg, nil
}

func LoadDoctorConfig() (*DoctorConfig, error) {
	var cfg DoctorConfig

	err := configor.New(&configor.Config{ErrorOnUnmatchedKeys: false}).Load(&cfg, "config.yml")
	if err != nil {
		return nil, errors.New("ERROR: Failed to load config file " + err.Error())
	}

	return &cfg, nil
}
