package services

import (
	"fmt"
	"slices"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vaaxooo/tgaccessbot/pkg/logger"
)

type SubscriptionService struct {
	logger        *logger.Service
	bot           *tgbotapi.BotAPI
	channels      []string
	AllowedStatus []string
}

// The NewSubscriptionService function creates a new SubscriptionService instance with the provided
// Telegram bot and list of channels.
func NewSubscriptionService(logger *logger.Service, bot *tgbotapi.BotAPI, channels []string) *SubscriptionService {
	return &SubscriptionService{
		logger:        logger,
		bot:           bot,
		channels:      channels,
		AllowedStatus: []string{"member", "administrator", "creator"},
	}
}

// This `IsUserSubscribed` method in the `SubscriptionService` struct is checking if a user with
// a specific `userID` is subscribed to all the channels listed in the `channels` slice.
func (s *SubscriptionService) IsUserSubscribed(userID int64) (bool, error) {
	for _, channel := range s.channels {
		member, err := s.bot.GetChatMember(tgbotapi.GetChatMemberConfig{
			ChatConfigWithUser: tgbotapi.ChatConfigWithUser{
				ChatID:             0, // Set appropriate ChatID if needed
				SuperGroupUsername: "@" + channel,
				UserID:             userID,
			},
		})
		if err != nil {
			s.logger.Error("❌ Error getting chat member: %v", err)

			return false, fmt.Errorf("error getting chat member: %w", err)
		}

		if !slices.Contains(s.AllowedStatus, member.Status) {
			return false, nil
		}

		continue
	}

	return true, nil
}
