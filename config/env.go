package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config struct contains the configuration of the application.
type Config struct {
	Debug              bool     `env:"DEBUG"                env-default:"false"`
	TelegramBotToken   string   `env:"TELEGRAM_BOT_TOKEN"   env-required:"true"`
	TelegramChannels   []string `env:"TELEGRAM_CHANNELS"    env-required:"true"`
	SuccessRedirectURL string   `env:"SUCCESS_REDIRECT_URL" env-required:"true"`
}

// MustLoadConfig loads the configuration from environment variables.
func MustLoadConfig() *Config {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(fmt.Errorf("configuration error: %w", err))
	}

	return &cfg
}
