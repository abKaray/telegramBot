package di

import (
	"go.uber.org/fx"
	"telegram_bot/pkg/chat_gpt"
	"telegram_bot/pkg/telegram_bot"
)

func AppProviders() []fx.Option {
	modules := []fx.Option{
		telegram_bot.Module,
		chat_gpt.Module,
	}

	return modules
}
