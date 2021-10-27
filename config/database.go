package config

type (
	Database struct {
		Mysql Mysql
	}

	Mysql struct {
		Host     string `env:"MYSQL_HOST"`
		Port     string `env:"MYSQL_PORT"`
		UserName string `env:"MYSQL_USERNAME"`
		Password string `env:"MYSQL_PASSWORD"`
		Database string `env:"MYSQL_DATABASE"`
	}
)
