package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"golara/config"
	"log"
)

type ConfigProvider struct{}

var cfg = config.Config{}

func (p *ConfigProvider) Boot() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n")
	}
}

func Config() config.Config {
	return cfg
}
