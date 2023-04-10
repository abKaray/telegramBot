package telegram_bot

import (
	"go.uber.org/fx"
	"os"
)

var Module = fx.Module("telegram_bot",
	fx.Provide(
		func() *TelegramBot {
			return &TelegramBot{ApiToken: os.Getenv("TELEGRAM_API_TOKEN")}
		},
		NewTelegramImp,
	),
)
