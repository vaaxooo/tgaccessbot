package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
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
	if err := godotenv.Load(); err != nil {
		panic(fmt.Errorf("error loading .env file: %w", err))
	}

	return &Config{
		Debug:              strings.ToLower(os.Getenv("DEBUG")) == "true",
		TelegramBotToken:   os.Getenv("TELEGRAM_BOT_TOKEN"),
		TelegramChannels:   strings.Split(os.Getenv("TELEGRAM_CHANNELS"), ","),
		SuccessRedirectURL: os.Getenv("SUCCESS_REDIRECT_URL"),
	}
}
