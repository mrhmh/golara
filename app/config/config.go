package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"golara/config"
	"log"
)

var cfg = config.Config{}

func InitConfig() {
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
