package app

// App is a structure that contains the main components of the application.
type App struct {
	Logger             Logger
	Bot                BotAPI
	Subscription       SubscriptionService
	LocaleStorage      LocaleProvider
	TelegramChannels   []string
	SuccessRedirectURL string
}

// NewApp creates a new instance of the App structure.
func NewApp(
	logger Logger,
	bot BotAPI,
	subService SubscriptionService,
	localeProvider LocaleProvider,
	channels []string,
	successURL string,
) *App {
	return &App{
		Logger:             logger,
		Bot:                bot,
		Subscription:       subService,
		LocaleStorage:      localeProvider,
		TelegramChannels:   channels,
		SuccessRedirectURL: successURL,
	}
}

// Start starts the application and listens for incoming updates.
func (a *App) Start() {
	a.Logger.Info("🚀 Bot successfully started")

	u := a.Bot.NewUpdate(0)
	u.Timeout = 60
	updates := a.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			a.HandleMessage(update)
		} else if update.CallbackQuery != nil {
			a.HandleCallbackQuery(update)
		}
	}
}

// func (a *App) toMessage(m *tgbotapi.Message) *Message {
// 	return &Message{
// 		ChatID:   m.Chat.ID,
// 		UserID:   m.From.ID,
// 		Text:     m.Text,
// 		Language: m.From.LanguageCode,
// 		Command:  m.Command(),
// 	}
// }

// func (a *App) toCallbackQuery(q *tgbotapi.CallbackQuery) *CallbackQuery {
// 	return &CallbackQuery{
// 		ChatID:   q.Message.Chat.ID,
// 		UserID:   q.From.ID,
// 		Data:     q.Data,
// 		QueryID:  q.ID,
// 		Language: q.From.LanguageCode,
// 	}
// }
