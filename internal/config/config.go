package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

type Config struct {
	Env      string `env:"ENV" `
	Logger   string `env:"LOGGER_LEVEL"`
	APP_PORT string `env:"APP_PORT"`
	CDN_HOST string `env:"CDN_HOST"`
}

func CreateConfig() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, errors.Wrap(err, "cannot read config from environment variables")
	}

	return &cfg, nil
}
