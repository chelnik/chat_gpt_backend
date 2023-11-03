package services

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"openai/internal/domain"
)

type Responder interface {
	ResponseSingle(query domain.SingleQuery) (openai.ChatCompletionResponse, error)
}

type SingleResponse struct {
	Responder
}

func NewSingleResponder() *SingleResponse {
	return &SingleResponse{}
}

// ResponseSingle стандартый ответ ChatGPT
func (r *SingleResponse) ResponseSingle(query domain.SingleQuery) (openai.ChatCompletionResponse, error) {
	client := openai.NewClient(query.Key)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: query.Inquiry,
				},
			},
		},
	)

	return resp, err
}
