package gamesgrpc

import (
	"github.com/ArtemLymarenko/steam-tg-bot/protos/gen/go/games"
	"github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/domain"
)

func mapDomainGameToGameResponse(game domain.Game) *games.Game {
	return &games.Game{
		Id:              game.Id,
		Name:            game.Name,
		Url:             game.GameInfo.Url,
		ImageUrl:        game.GameInfo.ImageUrl,
		InitialPrice:    game.GameInfo.InitialPrice,
		FinalPrice:      game.GameInfo.FinalPrice,
		DiscountPercent: game.GameInfo.DiscountPercent,
	}
}
