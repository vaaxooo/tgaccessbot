package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Message contains data about received message.
type Message struct {
	ChatID   int64
	UserID   int64
	Text     string
	Language string
	Command  string
}

// CallbackQuery contains data about received callback query.
type CallbackQuery struct {
	ChatID   int64
	UserID   int64
	Data     string
	QueryID  string
	Language string
}

// BotAPI — interface for Telegram bot API.
type BotAPI interface {
	Send(msg tgbotapi.Chattable) (tgbotapi.Message, error)
	Request(chattable tgbotapi.Chattable) (*tgbotapi.APIResponse, error)
	GetUpdatesChan(u tgbotapi.UpdateConfig) tgbotapi.UpdatesChannel
	NewUpdate(offset int) tgbotapi.UpdateConfig
	GetChatMember(chatID int64, channel string, userID int64) (tgbotapi.ChatMember, error)
}

// Logger — interface for logging.
type Logger interface {
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

// SubscriptionService — interface for checking user subscription status.
type SubscriptionService interface {
	IsUserSubscribed(userID int64) (bool, error)
}

// LocaleProvider — interface for getting localized messages.
type LocaleProvider interface {
	GetText(lang, key string) string
}
