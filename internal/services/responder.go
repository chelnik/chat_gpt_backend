package services

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"openai/internal/domain"
)

type Responder interface {
	ResponseSingle(query domain.SingleQuery) (openai.ChatCompletionResponse, error)
}

type ResponseService struct {
	Responder
}

func NewResponseService() *ResponseService {
	return &ResponseService{}
}

// ResponseSingle стандартый ответ ChatGPT
func (r *ResponseService) ResponseSingle(query domain.SingleQuery) (openai.ChatCompletionResponse, error) {
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
