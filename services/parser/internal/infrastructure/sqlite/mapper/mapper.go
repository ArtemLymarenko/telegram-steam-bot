package sqlitemap

import (
	"parser/internal/domain"
	"parser/internal/infrastructure/sqlite"
)

func GameEntityToCreateGameParams(game domain.Game) repository.CreateGameParams {
	return repository.CreateGameParams{
		ID:   game.Id,
		Name: game.Name,
	}
}

func GameInfoEntityToCreateGameInfoParams(gameInfo domain.GameInfo) repository.CreateGameInfoParams {
	return repository.CreateGameInfoParams{
		GameID:          gameInfo.GameId,
		ImageUrl:        toNullString(gameInfo.ImageUrl),
		InitialPrice:    toNullFloat64(gameInfo.InitialPrice),
		FinalPrice:      toNullFloat64(gameInfo.FinalPrice),
		DiscountPercent: toNullFloat64(gameInfo.DiscountPercent),
	}
}

func FindGameRowToGame(row repository.FindGameRow) domain.Game {
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
