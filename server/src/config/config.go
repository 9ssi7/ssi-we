package config

type App struct {
	Port string `env:"PORT" mapstructure:"PORT"`
	Host string `env:"HOST" mapstructure:"HOST"`
	Db   struct {
		Host       string `env:"MONGO_HOST" mapstructure:"MONGO_HOST"`
		Port       string `env:"MONGO_PORT" mapstructure:"MONGO_PORT"`
		Name       string `env:"MONGO_NAME" mapstructure:"MONGO_NAME"`
		Password   string `env:"MONGO_PASSWORD" mapstructure:"MONGO_PASSWORD"`
		Username   string `env:"MONGO_USERNAME" mapstructure:"MONGO_USERNAME"`
		Collection string `env:"MONGO_COLLECTION" mapstructure:"MONGO_COLLECTION"`
	}
	Cors struct {
		AllowOrigins     string `env:"CORS_ALLOW_ORIGINS" mapstructure:"CORS_ALLOW_ORIGINS"`
		AllowMethods     string `env:"CORS_ALLOW_METHODS" mapstructure:"CORS_ALLOW_METHODS"`
		AllowHeaders     string `env:"CORS_ALLOW_HEADERS" mapstructure:"CORS_ALLOW_HEADERS"`
		AllowCredentials bool   `env:"CORS_ALLOW_CREDENTIALS" mapstructure:"CORS_ALLOW_CREDENTIALS"`
	}
}
