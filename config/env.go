package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/vaaxooo/tgaccessbot/pkg/helpers"
)

// Config struct contains the configuration of the application
type Config struct {
	Debug              bool     `env:"DEBUG" env-default:"false"`
	TelegramBotToken   string   `env:"TELEGRAM_BOT_TOKEN" env-required:"true"`
	TelegramChannels   []string `env:"TELEGRAM_CHANNELS" env-required:"true" env-separator:","`
	SuccessRedirectURL string   `env:"SUCCESS_REDIRECT_URL" env-required:"true"`
}

// LoadConfig loads the configuration from the environment variables
func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	telegramChannels := helpers.SplitValue(os.Getenv("TELEGRAM_CHANNELS"), ",")

	return &Config{
		Debug:              strings.ToLower(os.Getenv("DEBUG")) == "true",
		TelegramBotToken:   os.Getenv("TELEGRAM_BOT_TOKEN"),
		TelegramChannels:   telegramChannels,
		SuccessRedirectURL: os.Getenv("SUCCESS_REDIRECT_URL"),
	}, nil
}
