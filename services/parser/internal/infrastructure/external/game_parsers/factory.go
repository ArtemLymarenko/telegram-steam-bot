package game_parsers

import (
	"context"
	"github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/domain/game"
)

type Parser interface {
	ParseGameInfo(ctx context.Context, gameId game.Id) game.Info
}

type ParserType int

const (
	SteamParser ParserType = iota
	EpicGamesParser
)

type ParserFactory struct {
	gameRepo GameRepository
}

func NewFactory(gameRepo GameRepository) *ParserFactory {
	return &ParserFactory{
		gameRepo: gameRepo,
	}
}

func (p *ParserFactory) CreateInstance(parserType ParserType) Parser {
	switch parserType {
	case SteamParser:
		return NewSteam(p.gameRepo)
	case EpicGamesParser:
		return nil
	}

	return nil
}
