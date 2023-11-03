package services

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"openai/internal/domain"
)

type Responder interface {
	ResponseSingle(query domain.SingleQuery) (domain.ChatResponse, error)
}

type ResponseService struct {
	Responder
}

func NewResponseService() *ResponseService {
	return &ResponseService{}
}

// ResponseSingle стандартый одиночный ответ ChatGPT
func (r *ResponseService) ResponseSingle(query domain.SingleQuery) (domain.ChatResponse, error) {
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

	return domain.Wrap(resp), err
}
