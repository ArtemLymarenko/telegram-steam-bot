package repository

import (
	"context"
	"database/sql"
	"errors"
	"parser/internal/domain"
)

type Games struct {
	queries *Queries
}

func NewGames(db DBTX) *Games {
	return &Games{queries: New(db)}
}

func (g *Games) WithTx(tx *sql.Tx) domain.GamesRepository {
	return &Games{
		queries: g.queries.WithTx(tx),
	}
}

func (g *Games) FindGame(ctx context.Context, gameId int64) (*domain.Game, error) {
	game, err := g.queries.findGame(ctx, gameId)
	if err != nil {
		return nil, errors.New("failed to find game")
	}

	found := findGameRowToGame(game)
	return &found, nil
}

func (g *Games) CreateGame(ctx context.Context, id int64, name string) error {
	return g.queries.createGame(ctx, createGameParams{
		ID:   id,
		Name: name,
	})
}

func (g *Games) CreateGameInfo(ctx context.Context, info domain.GameInfo) error {
	return g.queries.createGameInfo(ctx, gameInfoEntityToCreateGameInfoParams(info))
}

func (g *Games) FindUserGames(ctx context.Context, userId int64) ([]domain.Game, error) {
	games, err := g.queries.findUserGames(ctx, userId)
	if err != nil {
		return nil, errors.New("failed to find user games")
	}

	foundGames := findUserGamesRowsToGames(games)
	return foundGames, nil
}
