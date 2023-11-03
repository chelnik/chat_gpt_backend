package main

import (
	"log"
	"openai/configs"
	"openai/internal/client"
	"openai/internal/handlers"
	http_server "openai/internal/http-server"
	"openai/internal/services"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	// Read .env file.
	godotenv.Load()

	a := client.App{Key: os.Getenv("OPENAI_API_KEY")}
	a.Streaming()

	//StartHTTPServer()
}

func StartHTTPServer() {
	service := services.NewService()
	handler := handlers.NewHandler(*service)
	server := new(http_server.Server)

	serverCfg := configs.ServerConfig{Port: "8080"}

	go func() {
		log.Fatalf("Ошибка в запуске сервера - %v", server.Run(serverCfg, handler.Init()))
	}()
	log.Printf("Сервер успешно запустился на порту - %s", serverCfg.Port)

	// Простой graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
