package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vaaxooo/tgaccessbot/internal/locale"
	"github.com/vaaxooo/tgaccessbot/internal/services"
	"github.com/vaaxooo/tgaccessbot/pkg/logger"
)

// App is a structure that contains all the necessary services and settings for the application.
type App struct {
	Logger             *logger.Service
	Bot                *tgbotapi.BotAPI
	Subscription       *services.SubscriptionService
	LocaleStorage      *locale.Storage
	TelegramChannels   []string
	SuccessRedirectURL string
}
