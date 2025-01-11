package main

import (
	"github.com/joho/godotenv"
	"github.com/steam-tg-bot/services/bot/internal/infrastructure/telegram"
	v1Bot "github.com/steam-tg-bot/services/bot/internal/interface/bot"
	"github.com/steam-tg-bot/services/bot/internal/interface/bot/handlers"
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
