package gamesgrpc

import (
	"github.com/ArtemLymarenko/steam-tg-bot/protos/gen/go/games"
	"github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/domain/game"
)

func mapDomainGameToGameResponse(game game.Game) *games.Game {
	return &games.Game{
		Id:              int64(game.Id),
		Name:            string(game.Name),
		Url:             string(game.Info.Url),
		ImageUrl:        string(game.Info.ImageUrl),
		InitialPrice:    float64(game.Info.InitialPrice),
		FinalPrice:      float64(game.Info.FinalPrice),
		DiscountPercent: float64(game.Info.DiscountPercent),
	}
}
