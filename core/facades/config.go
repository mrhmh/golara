package facades

import (
	"golara/config"
)

var cfg = config.Config{}

func SetConfig(c config.Config) {
	cfg = c
}

func Config() config.Config {
	return cfg
}
