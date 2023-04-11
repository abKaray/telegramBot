package chat_service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/fx"
	"log"
	"telegram_bot/pkg/chat_gpt"
	"telegram_bot/pkg/telegram_bot"
)

func SendGreetingMessage(update tgbotapi.Update, bot *telegram_bot.TelegramBotImpl) bool {
	if update.Message.IsCommand() && update.Message.Command() == "start" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привіт, цей бот звязанний з ChatGPT, тому можеш сміливо писати все, що захочеш")
		_, err := bot.Bot.Send(msg)
		if err != nil {
			log.Println(err)
		}
		return true
	}

	return false
}

func SendMessage(lc fx.Lifecycle, bot *telegram_bot.TelegramBotImpl, chatGPT *chat_gpt.ChatGptImpl) {
	messages := make([]openai.ChatCompletionMessage, 0)

	bot.ListenForMessage(func(update tgbotapi.Update) {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: update.Message.Text,
		})

		isGreetingMessage := SendGreetingMessage(update, bot)

		if isGreetingMessage {
			return
		}

		update.Message.Text, _ = chatGPT.SendChatGPT(messages)

		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: update.Message.Text,
		})

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Bot.Send(msg)
	})
}
