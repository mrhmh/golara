package config

type App struct {
	Debug    bool `env:"APP_DEBUG"`
	RestPort uint `env:"APP_REST_PORT"`
	GrpcPORT uint `env:"APP_GRPC_PORT"`
}
