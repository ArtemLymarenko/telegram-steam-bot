package parserService

import (
	"context"
	"github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/domain/game"
	"log"
	"sync"
)

type ParserGameRepository interface {
	CreateGameInfo(ctx context.Context, info game.Info) error
}

type Parser interface {
	ParseAsync(concurrencyLevel int) <-chan game.Info
}

type ParserService struct {
	parserGameRepository ParserGameRepository
}

type ParserConfig struct {
	ReadWorkers  int
	WriteWorkers int
	Parser       Parser
}

func New(parserGameRepository ParserGameRepository) *ParserService {
	return &ParserService{
		parserGameRepository: parserGameRepository,
	}
}

func (s *ParserService) saveGamesInfoFromCh(gamesCh <-chan game.Info) {
	ctx := context.Background()
	for g := range gamesCh {
		err := s.parserGameRepository.CreateGameInfo(ctx, g)
		if err != nil {
			log.Println("Error while saving game info: ", err)
		}
	}
}

func (s *ParserService) ParseAndSaveAsync(readWorkers, writeWorkers int, parser Parser) {
	wg := sync.WaitGroup{}
	gamesCh := parser.ParseAsync(readWorkers)

	for range writeWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.saveGamesInfoFromCh(gamesCh)
		}()
	}

	wg.Wait()
}

func (s *ParserService) RunParsers(parser ...ParserConfig) {
	wg := sync.WaitGroup{}
	wg.Add(len(parser))
	for _, p := range parser {
		go func() {
			defer wg.Done()
			s.ParseAndSaveAsync(p.ReadWorkers, p.WriteWorkers, p.Parser)
		}()
	}
	wg.Wait()
}
