package chat_gpt

import (
	"go.uber.org/fx"
	"os"
)

var Module = fx.Module("chat_gpt",
	fx.Provide(
		func() *ChatGpt {
			return &ChatGpt{ApiToken: os.Getenv("CHAT_GPT_API_TOKEN")}
		},
		NewChatGPTImpl,
	),
)
