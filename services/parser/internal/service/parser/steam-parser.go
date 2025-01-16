package parserService

import (
	"github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/domain/game"
	"sync"
)

const batchLimit = 1024

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

func (steam *Steam) spawnWorkers(count int) {
	for i := 0; i < count; i++ {
		steam.wg.Add(1)
		go steam.worker()
	}
}

func (steam *Steam) worker() {
	defer steam.wg.Done()

	for id := range steam.inPool {
		//TODO: fetch data from steam by id

		steam.outBuffer <- game.Info{
			GameId: id,
		}
	}
}

func (steam *Steam) waitWorkerDone() {
	steam.wg.Wait()
	close(steam.outBuffer)
}

func (steam *Steam) ParseAsync(concurrencyLevel int) <-chan game.Info {
	steam.spawnWorkers(concurrencyLevel)
	go steam.waitWorkerDone()

	offset := 0
	for {
		gameIds := steam.gamesRepo.GetGameIdsWithOffset(offset, batchLimit)
		if len(gameIds) == 0 {
			break
		}

		for _, id := range gameIds {
			steam.inPool <- id
		}

		offset += batchLimit
	}
	close(steam.inPool)

	return steam.outBuffer
}
