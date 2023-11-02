package main

import (
	"openai/internal/client"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Read .env file.
	godotenv.Load()

	a := client.App{Key: os.Getenv("OPENAI_API_KEY")}
	a.Streaming()
}
