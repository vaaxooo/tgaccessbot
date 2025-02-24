package services

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vaaxooo/tgaccessbot/pkg/helpers"
	"github.com/vaaxooo/tgaccessbot/pkg/logger"
)

var ALLOWED_STATUS = []string{"member", "administrator", "creator"}

type SubscriptionService struct {
	bot      *tgbotapi.BotAPI
	channels []string
}

// The NewSubscriptionService function creates a new SubscriptionService instance with the provided
// Telegram bot and list of channels.
func NewSubscriptionService(bot *tgbotapi.BotAPI, channels []string) *SubscriptionService {
	return &SubscriptionService{
		bot:      bot,
		channels: channels,
	}
}

// This `IsUserSubscribed` method in the `SubscriptionService` struct is checking if a user with a
// specific `userID` is subscribed to all the channels listed in the `channels` slice. Here's a
// breakdown of what the method does:
func (s *SubscriptionService) IsUserSubscribed(userID int64) (bool, error) {
	for _, channel := range s.channels {
		member, err := s.bot.GetChatMember(tgbotapi.GetChatMemberConfig{
			ChatConfigWithUser: tgbotapi.ChatConfigWithUser{
				SuperGroupUsername: "@" + channel,
				UserID:             userID,
			},
		})
		if err != nil {
			logger.Error("❌ Error getting chat member: %v", err)
			return false, err
		}
		if !helpers.Contains(ALLOWED_STATUS, member.Status) {
			return false, nil
		}
		continue
	}
	return true, nil
}
