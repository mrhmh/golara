package config

type (
	Config struct {
		Database Database
	}

	Database struct {
		Mysql Mysql
	}

	Mysql struct {
		Host string `env:"HOST"`
	}
)
