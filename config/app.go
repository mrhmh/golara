package config

type App struct {
	RestPort uint `env:"APP_REST_PORT"`
	GrpcPORT uint `env:"APP_GRPC_PORT"`
}
