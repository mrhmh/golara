package app

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"golara/config"
	"log"
)

var Config = config.Config{}

func configServiceProvider() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
	if err := env.Parse(&Config); err != nil {
		log.Fatalf("%+v\n")
	}
}
