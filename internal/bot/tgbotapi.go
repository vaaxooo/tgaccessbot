package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// TelegramBot — interface for Telegram bot API.
type TelegramBot struct {
	API *tgbotapi.BotAPI
}

// NewTelegramBot creates a new TelegramBot instance with the provided Telegram bot API.
func NewTelegramBot(api *tgbotapi.BotAPI) *TelegramBot {
	return &TelegramBot{API: api}
}

// Send sends a message to the chat.
func (t *TelegramBot) Send(msg tgbotapi.Chattable) (tgbotapi.Message, error) {
	message, err := t.API.Send(msg)
	if err != nil {
		return tgbotapi.Message{}, fmt.Errorf("failed to send message: %w", err)
	}

	return message, nil
}

// Request sends a request to the Telegram bot API.
func (t *TelegramBot) Request(chattable tgbotapi.Chattable) (*tgbotapi.APIResponse, error) {
	resp, err := t.API.Request(chattable)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	return resp, nil
}

// GetUpdatesChan returns a channel for receiving updates.
func (t *TelegramBot) GetUpdatesChan(u tgbotapi.UpdateConfig) tgbotapi.UpdatesChannel {
	return t.API.GetUpdatesChan(u)
}

// NewUpdate creates a new update configuration with the provided offset.
func (t *TelegramBot) NewUpdate(offset int) tgbotapi.UpdateConfig {
	return tgbotapi.NewUpdate(offset)
}

// GetChatMember returns information about a chat member.
func (t *TelegramBot) GetChatMember(chatID int64, channel string, userID int64) (tgbotapi.ChatMember, error) {
	memberConfig := tgbotapi.GetChatMemberConfig{
		ChatConfigWithUser: tgbotapi.ChatConfigWithUser{
			ChatID:             chatID,
			SuperGroupUsername: channel,
			UserID:             userID,
		},
	}

	resp, err := t.API.GetChatMember(memberConfig)
	if err != nil {
		return tgbotapi.ChatMember{}, fmt.Errorf("failed to get chat member: %w", err)
	}

	return resp, nil
}
