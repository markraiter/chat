package configs

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v8"
)

type Server struct {
	AppAddress string `env:"APP_PORT" envDefault:"9000"`
	// AppReadTimeout  time.Duration `env:"APP_READ_TIMEOUT" envDefault:"60s"`
	// AppWriteTimeout time.Duration `env:"APP_WRITE_TIMEOUT" envDefault:"60s"`
	// AppIdleTimeout  time.Duration `env:"APP_IDLE_TIMEOUT" envDefault:"60s"`
}

type Postgres struct {
	Driver     string `env:"POSTGRES_DRIVER" envDefault:"postgres"`
	ConnString string `env:"POSTGRES_CONN_STRING" envDefault:"postgresql://root:password@localhost:5433/chat-app-db?sslmode=disable"`
}

// type SMTP struct {
// 	MailAccount       string `env:"SMTP_ACCOUNT"`
// 	AccountPassword   string `env:"SMTP_PASSWORD"`
// 	SMTPServerAddress string `env:"SMTP_ADDRESS"`
// 	SMTPPort          string `env:"SMTP_PORT"`
// }

type AuthConfig struct {
	Salt            string        `env:"APP_SALT,notEmpty"`
	SigningKey      string        `env:"SIGNING_KEY,notEmpty"`
	AccessTokenTTL  time.Duration `env:"ACCESS_TOKEN_TTL" envDefault:"15m"`
	RefreshTokenTTL time.Duration `env:"REFRESH_TOKEN_TTL" envDefault:"24h"`
}

type Config struct {
	Server   Server
	DB       Postgres
	Auth     AuthConfig
	LogLevel string `env:"LOG_LEVEL" envDefault:"INFO"`
	BaseURL  string `env:"BASE_URL" envDefault:"https://example.com"`
	// SMTP     SMTP
}

func InitConfig() (Config, error) {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return Config{}, fmt.Errorf("error while parsing .env: %w", err)
	}

	cfg.Server.AppAddress = ":" + cfg.Server.AppAddress

	return cfg, nil
}
