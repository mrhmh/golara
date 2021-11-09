package providers

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"golara/config"
	"golara/core/facades"
	"log"
)

type ConfigProvider struct{}

func (p *ConfigProvider) Boot() {

	var cfg = config.Config{}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n")
	}

	facades.SetConfig(cfg)
}
