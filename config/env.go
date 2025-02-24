package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/vaaxooo/tgaccessbot/pkg/helpers"
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

	telegramChannels := helpers.SplitString(os.Getenv("TELEGRAM_CHANNELS"), ",")

	return &Config{
		Debug:              os.Getenv("DEBUG") == "true",
		TelegramBotToken:   os.Getenv("TELEGRAM_BOT_TOKEN"),
		TelegramChannels:   telegramChannels,
		SuccessRedirectURL: os.Getenv("SUCCESS_REDIRECT_URL"),
	}
}
