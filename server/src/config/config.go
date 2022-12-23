package config

type App struct {
	Port string `env:"PORT"`
	Host string `env:"HOST"`
	Db   struct {
		Host       string `env:"MONGO_HOST"`
		Port       string `env:"MONGO_PORT"`
		Name       string `env:"MONGO_NAME"`
		Password   string `env:"MONGO_PASSWORD"`
		Username   string `env:"MONGO_USERNAME"`
		Collection string `env:"MONGO_COLLECTION"`
	}
	Cors struct {
		AllowOrigins     string `env:"CORS_ALLOW_ORIGINS"`
		AllowMethods     string `env:"CORS_ALLOW_METHODS"`
		AllowHeaders     string `env:"CORS_ALLOW_HEADERS"`
		AllowCredentials bool   `env:"CORS_ALLOW_CREDENTIALS"`
	}
}
