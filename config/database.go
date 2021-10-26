package config

type (
	Database struct {
		Mysql Mysql
	}

	Mysql struct {
		Host string `env:"HOST"`
	}
)
