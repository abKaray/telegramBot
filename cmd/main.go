package main

import (
	"telegram_bot/di"
	"telegram_bot/modules/telegram_bot"
)

func main() {
	fxOptions := di.AppProviders()

	telegram_bot.StartTelegram(fxOptions)
}
