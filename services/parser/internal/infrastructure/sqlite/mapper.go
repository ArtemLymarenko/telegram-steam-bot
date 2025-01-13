package repository

import (
	"github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/domain/game"
)

func gameInfoEntityToCreateGameInfoParams(gameInfo game.Info) createGameInfoParams {
	return createGameInfoParams{
		GameID:          int64(gameInfo.GameId),
		Url:             toNullString(string(gameInfo.Url)),
		ImageUrl:        toNullString(string(gameInfo.ImageUrl)),
		InitialPrice:    toNullFloat64(float64(gameInfo.InitialPrice)),
		FinalPrice:      toNullFloat64(float64(gameInfo.FinalPrice)),
		DiscountPercent: toNullFloat64(float64(gameInfo.DiscountPercent)),
	}
}

func findGameRowToGame(row findGameRow) game.Game {
	return game.Game{
		Id:   game.Id(row.Game.ID),
		Name: game.Name(row.Game.Name),
		Info: game.Info{
			GameId:          game.Id(row.Game.ID),
			Url:             game.Url(toString(row.GameInfo.Url)),
			ImageUrl:        game.ImageUrl(toString(row.GameInfo.ImageUrl)),
			InitialPrice:    game.Price(toFloat64(row.GameInfo.InitialPrice)),
			FinalPrice:      game.Price(toFloat64(row.GameInfo.FinalPrice)),
			DiscountPercent: game.DiscountPercent(toFloat64(row.GameInfo.DiscountPercent)),
		},
	}
}

func findUserGamesRowsToGames(rows []findUserGamesRow) []game.Game {
	result := make([]game.Game, len(rows))
	for i, row := range rows {
		result[i] = findGameRowToGame(findGameRow(row))
	}
	return result
}
