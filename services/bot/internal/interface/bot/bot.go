package v1Bot

import (
	"bot/internal/infrastructure/telegram"
	"bot/internal/interface/bot/handlers"
)

func GetRouter(handlers *handlers.BotHandlers) *telegram.Router {
	router := telegram.NewRouter()

	router.UseMiddleware(handlers.HelloMiddleware)

	router.AddInlineQuery(handlers.InlineEchoQuery)

	router.AddHandler(telegram.HandlerTypeCmd, "open", handlers.Open)
	router.AddHandler(telegram.HandlerTypeCmd, "close", handlers.Close)
	router.AddHandler(telegram.HandlerTypeCmd, "help", handlers.Help)
	router.AddHandler(telegram.HandlerTypeCmd, "add_steam_game", handlers.AddGame)
	router.AddHandler(telegram.HandlerTypeCmd, "check_my_games", handlers.CheckMyGames)

	return router
}
