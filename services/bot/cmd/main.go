package main

import (
	"github.com/ArtemLymarenko/steam-tg-bot/services/bot/internal/infrastructure/telegram"
	v1Bot "github.com/ArtemLymarenko/steam-tg-bot/services/bot/internal/interface/bot"
	"github.com/ArtemLymarenko/steam-tg-bot/services/bot/internal/interface/bot/handlers"
	gamesgrpc "github.com/ArtemLymarenko/steam-tg-bot/services/bot/internal/interface/grpc/games"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token := os.Getenv("TELEGRAM_TOKEN")

	grpcConn, err := grpc.NewClient(
		"localhost:44044",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer grpcConn.Close()

	api := gamesgrpc.NewClientApi(grpcConn)

	botHandlers := handlers.NewBotHandlers(api)
	routes := v1Bot.GetRouter(botHandlers)

	telegramBot := telegram.NewBot(token, true, routes)
	telegramBot.Listen()
}
