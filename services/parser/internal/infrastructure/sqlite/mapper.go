package repository

import (
	"parser/internal/domain"
)

func gameInfoEntityToCreateGameInfoParams(gameInfo domain.GameInfo) createGameInfoParams {
	return createGameInfoParams{
		GameID:          gameInfo.GameId,
		ImageUrl:        toNullString(gameInfo.ImageUrl),
		InitialPrice:    toNullFloat64(gameInfo.InitialPrice),
		FinalPrice:      toNullFloat64(gameInfo.FinalPrice),
		DiscountPercent: toNullFloat64(gameInfo.DiscountPercent),
	}
}

func findGameRowToGame(row findGameRow) domain.Game {
	return domain.Game{
		Id:   row.Game.ID,
		Name: row.Game.Name,
		GameInfo: domain.GameInfo{
			GameId:          row.Game.ID,
			ImageUrl:        toString(row.GameInfo.ImageUrl),
			InitialPrice:    toFloat64(row.GameInfo.InitialPrice),
			FinalPrice:      toFloat64(row.GameInfo.FinalPrice),
			DiscountPercent: toFloat64(row.GameInfo.DiscountPercent),
		},
	}
}

func findUserGamesRowToGame(row findUserGamesRow) domain.Game {
	return domain.Game{
		Id:   row.GameID,
		Name: row.GameName,
		GameInfo: domain.GameInfo{
			GameId:          row.GameID,
			ImageUrl:        toString(row.ImageUrl),
			InitialPrice:    toFloat64(row.InitialPrice),
			FinalPrice:      toFloat64(row.FinalPrice),
			DiscountPercent: toFloat64(row.DiscountPercent),
		},
	}
}

func findUserGamesRowsToGames(rows []findUserGamesRow) []domain.Game {
	result := make([]domain.Game, len(rows))
	for i, row := range rows {
		result[i] = findUserGamesRowToGame(row)
	}
	return result
}
