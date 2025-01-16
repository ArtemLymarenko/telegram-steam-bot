package parserService

import "github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/domain/game"

type PaserType int

const (
	SteamParserType PaserType = iota
	EpicGamesParserType
)

type GameRepository interface {
	GetGameIdsWithOffset(offset, limit int) []game.Id
}

type ParserFactory struct {
	gameRepo GameRepository
}

func NewParserFactory(gameRepo GameRepository) *ParserFactory {
	return &ParserFactory{
		gameRepo: gameRepo,
	}
}

func (p *ParserFactory) CreateInstance(parserType PaserType) Parser {
	switch parserType {
	case SteamParserType:
		return NewSteam(p.gameRepo)
	case EpicGamesParserType:
		return nil
	}

	return nil
}
