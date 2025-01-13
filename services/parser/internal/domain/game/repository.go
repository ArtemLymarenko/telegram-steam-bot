package game

import (
	"context"
	"database/sql"
)

type Repository interface {
	FindGame(ctx context.Context, gameId Id) (Game, error)
	CreateGame(ctx context.Context, id Id, name Name) error
	CreateGameInfo(ctx context.Context, info Info) error
	FindUserGames(ctx context.Context, userId UserId) ([]Game, error)
	AddUserGame(ctx context.Context, userId UserId, gameId Id) error
	DeleteGameById(ctx context.Context, gameId Id) error
	SearchGamesByName(ctx context.Context, name Name) ([]Game, error)
	WithTx(tx *sql.Tx) Repository
}
