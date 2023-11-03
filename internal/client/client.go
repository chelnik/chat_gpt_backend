package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"io"
)

// App структура в которой храним ключ
type App struct {
	Key string
}

// Streaming streaming потоковый ответ от чат гпт
// пишет посимвольно ответ от чат гпт в консоль и в файл
func (a App) Streaming() {
	c := openai.NewClient(a.Key)
	ctx := context.Background()
	reqs := readFromFiles()
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 1000,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: reqs[0],
				Name:    "User",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: reqs[1],
				Name:    "Assistant",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: reqs[2],
				Name:    "User",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: reqs[3],
				Name:    "Assistant",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: reqs[4],
				Name:    "User",
			},
		},
		Stream: true,
	}
	stream, err := c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	fmt.Printf("Stream response: ")
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			return
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return
		}

		fmt.Printf(response.Choices[0].Delta.Content)
		content := fmt.Sprintf("%s", response.Choices[0].Delta.Content)
		err = writeToFile("answer.txt", content)
	}
}

// ChatResponse chatResponse не потоковый ответ от чат гпт
// отдает сразу же ответ от чат гпт (не печатая его посимвольно)
func (a App) ChatResponse() {
	client := openai.NewClient(a.Key)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "сгенерируй название для изображения в формате picture_<uid>.png",
					// Content: "напиши сервер на гоу",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
