package domain

import "github.com/sashabaranov/go-openai"

type ChatResponse struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int64                  `json:"created"`
	Model   string                 `json:"model"`
	Choices []ChatCompletionChoice `json:"choices"`
	Usage   Usage                  `json:"usage"`
}

type ChatCompletionChoice struct {
	Index        int                   `json:"index"`
	Message      ChatCompletionMessage `json:"message"`
	FinishReason string                `json:"finish_reason"`
}
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
type ChatCompletionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Name    string `json:"name,omitempty"`
}

func Wrap(resp openai.ChatCompletionResponse) ChatResponse {
	response := ChatResponse{
		ID:      resp.ID,
		Object:  resp.Object,
		Created: resp.Created,
		Model:   resp.Model,
		Choices: []ChatCompletionChoice{
			{
				Index: resp.Choices[0].Index,
				Message: ChatCompletionMessage{
					Role:    resp.Choices[0].Message.Role,
					Content: resp.Choices[0].Message.Content,
					Name:    resp.Choices[0].Message.Name,
				},
				FinishReason: resp.Choices[0].FinishReason,
			},
		},
		Usage: Usage{
			PromptTokens:     resp.Usage.PromptTokens,
			CompletionTokens: resp.Usage.CompletionTokens,
			TotalTokens:      resp.Usage.TotalTokens,
		},
	}

	return response
}
