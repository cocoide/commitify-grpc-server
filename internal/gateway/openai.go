package gateway

import (
	"context"
	"os"

	"github.com/sashabaranov/go-openai"
)

//go:generate mockgen -source=openai.go -destination=../../mock/openai.go
type OpenAIGateway interface {
	GetAnswerFromPrompt(prompt string) (string, error)
}

type openAIGateway struct {
	client *openai.Client
	ctx    context.Context
}

func NewOpenAIGateway(ctx context.Context) OpenAIGateway {
	client := openai.NewClient(os.Getenv("GPT_API_KEY"))
	return &openAIGateway{client: client, ctx: ctx}
}

func (og *openAIGateway) GetAnswerFromPrompt(prompt string) (string, error) {
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		Temperature: 0.01,
	}
	res, err := og.client.CreateChatCompletion(og.ctx, req)
	if err != nil {
		return "", err
	}
	answer := res.Choices[0].Message.Content
	return answer, nil
}
