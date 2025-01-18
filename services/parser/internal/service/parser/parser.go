package parser_service

import (
	"context"
	"github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/domain/game"
	"github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/infrastructure/external/game_parsers"
	"github.com/ArtemLymarenko/steam-tg-bot/services/parser/pkg/thread_pool"
	"sync"
)

type GameRepository interface {
	CreateGameInfo(ctx context.Context, info game.Info) error
	GetGameIdsWithoutInfoBatched(offset, limit int) ([]game.Id, int)
	GetCountRows(ctx context.Context) int64
}

type ParserHandler struct {
	gamesRepo GameRepository
}

func New(gamesRepo GameRepository) *ParserHandler {
	return &ParserHandler{
		gamesRepo: gamesRepo,
	}
}

type ParserConfig struct {
	Workers int
	Parser  game_parsers.Parser
}

func (s *ParserHandler) ParseAndSaveInfoTask(id game.Id, parser game_parsers.Parser) thread_pool.TaskFunc {
	return func() error {
		ctx := context.Background()
		gameInfo := parser.ParseGameInfo(ctx, id)
		err := s.gamesRepo.CreateGameInfo(ctx, gameInfo)
		return err
	}
}

func (s *ParserHandler) ParseGameInfoAsync(config *ParserConfig) {
	const batchLimit = 1024

	pool := thread_pool.New(batchLimit)
	defer pool.TerminateWait()

	pool.RunWorkers(config.Workers)

	totalRows := s.gamesRepo.GetCountRows(context.Background())
	offset := 0
	for offset <= int(totalRows) {
		gameIds, actualOffset := s.gamesRepo.GetGameIdsWithoutInfoBatched(offset, batchLimit)

		for _, id := range gameIds {
			task := s.ParseAndSaveInfoTask(id, config.Parser)
			pool.AddTask(task)
		}

		offset += actualOffset
	}
}

func (s *ParserHandler) RunParsersAsync(run func(config *ParserConfig), parsersCfgs ...*ParserConfig) {
	wg := sync.WaitGroup{}
	wg.Add(len(parsersCfgs))

	for _, cfg := range parsersCfgs {
		go func() {
			defer wg.Done()
			run(cfg)
		}()
	}

	wg.Wait()
}
