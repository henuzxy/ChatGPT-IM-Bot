package openaiProxy

import (
	"chatgpt-im-bot/config"
	"chatgpt-im-bot/logger"
	"context"
	"errors"
	"github.com/sashabaranov/go-openai"
)

var globalClinet *openai.Client

func RegisterOpenAI() {
	globalClinet = openai.NewClient(config.GetConfig().OpenAIConfig.OpenAIKey)
}

func GetChatGPTSimpleReply(ctx context.Context, msg string) (string, error) {
	resp, err := globalClinet.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: msg,
				},
			},
		})
	if err != nil {
		logger.Errorf("chatGPT reply error:%s", err)
		return "", errors.New("chatGPT error")
	}
	if len(resp.Choices) < 1 {
		return "", errors.New("chatGPT response empty")
	}
	return resp.Choices[0].Message.Content, nil
}

func GetOpenAIImage(ctx context.Context, msg string) (string, error) {
	return "", nil
}
