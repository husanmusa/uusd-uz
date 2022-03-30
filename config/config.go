package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment       string
	PostgresHost      string
	PostgresPort      int
	PostgresDatabase  string
	PostgresUser      string
	PostgresPassword  string
	LogLevel          string
	ReviewServiceHost string
	ReviewServicePort int
	HTTPPort          string
}

func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "uusduz"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "husanmsua"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "pass"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "release"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
