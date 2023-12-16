package gateway

import (
	"context"
	"github.com/cocoide/commitify-grpc-server/internal/domain/service"
	"os"

	"github.com/sashabaranov/go-openai"
)

type openAIGateway struct {
	client *openai.Client
	ctx    context.Context
}

func NewOpenAIGateway(ctx context.Context) service.NLPService {
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
