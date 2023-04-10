package telegram_bot

import (
	"go.uber.org/fx"
	chat_service "telegram_bot/modules/telegram_bot/internal/service"
)

func StartTelegram(fxOptions []fx.Option) {

	fxOptions = append(
		fxOptions,
		fx.Invoke(
			chat_service.SendMessage,
		),
	)

	app := fx.New(fxOptions...)

	app.Run()
}
