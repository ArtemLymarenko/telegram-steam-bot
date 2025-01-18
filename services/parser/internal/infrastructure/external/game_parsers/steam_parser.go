package game_parsers

import (
	"context"
	"github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/domain/game"
	"sync"
)

const batchLimit = 1024

type GameRepository interface {
	GetGameIdsWithoutInfoBatched(offset, limit int) []game.Id
}

// Steam is a class that implements Parser interface
// with concurrencyLevel by default is 1.
type Steam struct {
	gamesRepo GameRepository
	inPool    chan game.Id
	outBuffer chan game.Info
	wg        *sync.WaitGroup
}

func NewSteam(gamesRepo GameRepository) *Steam {
	return &Steam{
		gamesRepo: gamesRepo,
		inPool:    make(chan game.Id, batchLimit),
		outBuffer: make(chan game.Info),
		wg:        &sync.WaitGroup{},
	}
}

func (steam *Steam) ParseGameInfo(ctx context.Context, gameId game.Id) game.Info {
	return game.Info{}
}
