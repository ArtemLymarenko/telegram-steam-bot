package main

import (
	"bot/internal/infrastructure/telegram"
	v1Bot "bot/internal/interface/bot"
	"bot/internal/interface/bot/handlers"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token := os.Getenv("TELEGRAM_TOKEN")

	botHandlers := handlers.NewBotHandlers()
	routes := v1Bot.GetRouter(botHandlers)

	telegramBot := telegram.NewBot(token, true, routes)
	telegramBot.Listen()
}
