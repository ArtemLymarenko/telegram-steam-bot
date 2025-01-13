package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/domain/game"
)

type Games struct {
	queries *Queries
}

func NewGames(db DBTX) *Games {
	return &Games{queries: New(db)}
}

func (g *Games) WithTx(tx *sql.Tx) game.Repository {
	return &Games{
		queries: g.queries.WithTx(tx),
	}
}

func (g *Games) FindGame(ctx context.Context, gameId game.Id) (game.Game, error) {
	found, err := g.queries.findGame(ctx, int64(gameId))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return game.Game{}, errors.New("failed to find game")
		}
		return game.Game{}, err
	}

	mappedGame := findGameRowToGame(found)
	return mappedGame, nil
}

func (g *Games) CreateGame(ctx context.Context, id game.Id, name game.Name) error {
	return g.queries.createGame(ctx, createGameParams{
		ID:   int64(id),
		Name: string(name),
	})
}

func (g *Games) CreateGameInfo(ctx context.Context, info game.Info) error {
	return g.queries.createGameInfo(ctx, gameInfoEntityToCreateGameInfoParams(info))
}

func (g *Games) FindUserGames(ctx context.Context, userId game.UserId) ([]game.Game, error) {
	games, err := g.queries.findUserGames(ctx, int64(userId))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("failed to find user games")
		}
		return nil, err
	}

	foundGames := findUserGamesRowsToGames(games)
	return foundGames, nil
}

func (g *Games) AddUserGame(ctx context.Context, userId game.UserId, gameId game.Id) error {
	err := g.queries.addUserGame(ctx, addUserGameParams{
		UserID: int64(userId),
		GameID: int64(gameId),
	})
	return err
}

func (g *Games) DeleteGameById(ctx context.Context, gameId game.Id) error {
	err := g.queries.deleteGameById(ctx, int64(gameId))
	return err
}

func (g *Games) SearchGamesByName(ctx context.Context, name game.Name) ([]game.Game, error) {
	found, err := g.queries.searchGamesByName(ctx, string(name))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("failed to search games by name")
		}
		return nil, err
	}

	games := searchUserGamesRowsToGames(found)
	return games, err
}
