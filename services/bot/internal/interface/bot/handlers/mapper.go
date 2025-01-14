package handlers

import (
	"fmt"
	"github.com/ArtemLymarenko/steam-tg-bot/services/bot/internal/infrastructure/telegram"
	"github.com/ArtemLymarenko/steam-tg-bot/services/bot/internal/interface/dto"
)

func prepareGameDescription(initialPrice, finalPrice, discountPercent float64) string {
	return fmt.Sprintf("Initial price: %.2f USD\nFinal price: %.2f USD\nDiscount: %.2f%\n", initialPrice, finalPrice, discountPercent)
}

func prepareAddGameCmd(gameId int64) string {
	return fmt.Sprintf("/add_steam_game %d", gameId)
}

func prepareArticlesForGameSearch(games []dto.Game) []telegram.Article {
	articles := make([]telegram.Article, len(games))

	for i, game := range games {
		article := telegram.Article{
			Url:        game.ImageUrl,
			Title:      game.Name,
			Desc:       prepareGameDescription(game.InitialPrice, game.FinalPrice, game.DiscountPercent),
			TextToSend: prepareAddGameCmd(game.Id),
		}
		articles[i] = article
	}

	return articles
}
