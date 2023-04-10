package chat_gpt

import (
	"context"
	openai "github.com/sashabaranov/go-openai"
)

type ChatGptImpl struct {
	ChatGpt *openai.Client
}

func NewChatGPTImpl(cfg *ChatGpt) *ChatGptImpl {
	chatGPT := openai.NewClient(cfg.ApiToken)

	return &ChatGptImpl{ChatGpt: chatGPT}
}

func (c *ChatGptImpl) SendChatGPT(messages []openai.ChatCompletionMessage) (string, error) {
	resp, err := c.ChatGpt.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: messages,
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
