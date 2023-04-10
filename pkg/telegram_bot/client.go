package telegram_bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const (
	defaultTimeoutMilliseconds = 60
)

type TelegramBotImpl struct {
	Bot *tgbotapi.BotAPI
}

func NewTelegramImp(cfg *TelegramBot) *TelegramBotImpl {
	bot, err := tgbotapi.NewBotAPI(cfg.ApiToken)

	if err != nil {
		log.Panic(err)
	}

	return &TelegramBotImpl{Bot: bot}
}

func (t *TelegramBotImpl) ListenForMessage(fn func(message tgbotapi.Update)) {
	u := tgbotapi.UpdateConfig{
		Offset:  0,
		Limit:   0,
		Timeout: defaultTimeoutMilliseconds,
	}

	updates := t.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		fn(update)
	}
}
