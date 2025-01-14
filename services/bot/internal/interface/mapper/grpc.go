package mapper

import (
	"github.com/ArtemLymarenko/steam-tg-bot/protos/gen/go/games"
	"github.com/ArtemLymarenko/steam-tg-bot/services/bot/internal/interface/dto"
)

func GameGrpcToDto(game *games.Game) dto.Game {
	return dto.Game{
		Id:              game.Id,
		Name:            game.Name,
		Url:             game.Url,
		ImageUrl:        game.ImageUrl,
		InitialPrice:    game.InitialPrice,
		FinalPrice:      game.FinalPrice,
		DiscountPercent: game.DiscountPercent,
	}
}

func GamesGrpcToDto(game []*games.Game) []dto.Game {
	res := make([]dto.Game, 0, len(game))
	for _, g := range game {
		res = append(res, GameGrpcToDto(g))
	}
	return res
}
