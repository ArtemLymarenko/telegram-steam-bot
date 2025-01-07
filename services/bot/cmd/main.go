package main

import (
	"bot/internal/infrastructure/telegram"
	v1Bot "bot/internal/interface/bot"
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

	router := v1Bot.GetRouter()
	newBot := telegram.NewBot(token, true, router)
	newBot.Listen()
}
