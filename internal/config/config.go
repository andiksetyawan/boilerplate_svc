package config

import (
	"github.com/andiksetyawan/config"
)

type Config struct {
	ServiceName string `env:"SERVICE_NAME,required"`
	ServicePort int    `env:"SERVICE_PORT" envDefault:"8080"`
	ServiceEnv  string `env:"SERVICE_ENV" envDefault:"development"`

	DBName     string `env:"DB_NAME"`
	DBUsername string `env:"DB_USERNAME"`
	DBPassword string `env:"DB_PASSWORD"`
	DBHost     string `env:"DB_HOST"`
	DBPort     string `env:"DB_PORT"`
	//DBSSLMode           string `env:"DB_SSL_MODE"`
	//DBConnectionTimeout int    `env:"DB_CONNECTION_TIMEOUT"`
	//DBMaxIdleConns      int    `env:"DB_MAX_IDLE_CONNS"`
	//DBMaxOpenConns      int    `env:"DB_MAX_OPEN_CONNS"`
	//DBMaxLifetimeMinute int    `env:"DB_MAX_LIFETIME_MINUTE"`

	JwtSigningKey string `env:"JWT_SIGNING_KEY"`

	OtelJaegerEndpoint string `env:"OTEL_JAEGER_ENDPOINT" envDefault:"http://localhost:14268/api/traces"`
}

func NewConfig() Config {
	configuration := Config{}
	if err := config.New().Load(&configuration); err != nil {
		panic(err)
	}

	return configuration
}
