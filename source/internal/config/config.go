package config

import "log/slog"

type Config struct {
	DBHost     string `env:"DB_HOST"     env-default:"localhost"`
	DBUser     string `env:"DB_USER"     env-default:"postgres"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`
	DBPort     string `env:"DB_PORT"     env-default:"5432"`

	ListenAddr string `env:"LISTEN_ADDR" env-default:"0.0.0.0"`
	ListenPort uint16 `env:"LISTEN_PORT" env-default:"1243"`

	LogLevel  slog.Level `env:"OCTOPUS_LOG_LEVEL"  env-default:"INFO"`
	LogFormat string     `env:"OCTOPUS_LOG_FORMAT" env-default:"json"`

	// AuthPublicKey string `env:"OCTOPUS_AUTH_PUBLIC_KEY" env-default:""`
}
