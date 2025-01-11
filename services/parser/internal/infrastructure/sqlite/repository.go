package repository

import (
	"context"
	"errors"
	"parser/internal/domain"
)

func (q *Queries) FindGame(ctx context.Context, gameId int64) (*domain.Game, error) {
	game, err := q.findGame(ctx, gameId)
	if err != nil {
		return nil, errors.New("failed to find game")
	}

	found := findGameRowToGame(game)
	return &found, nil
}

func (q *Queries) CreateGame(ctx context.Context, id int64, name string) error {
	return q.createGame(ctx, createGameParams{
		ID:   id,
		Name: name,
	})
}

func (q *Queries) CreateGameInfo(ctx context.Context, info domain.GameInfo) error {
	return q.createGameInfo(ctx, gameInfoEntityToCreateGameInfoParams(info))
}

func (q *Queries) FindUserGames(ctx context.Context, userId int64) ([]domain.Game, error) {
	games, err := q.findUserGames(ctx, userId)
	if err != nil {
		return nil, errors.New("failed to find user games")
	}

	foundGames := findUserGamesRowsToGames(games)
	return foundGames, nil
}
