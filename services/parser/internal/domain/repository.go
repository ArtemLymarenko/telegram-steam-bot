package domain

import (
	"context"
	"database/sql"
)

type GamesRepository interface {
	FindGame(ctx context.Context, gameId int64) (*Game, error)
	CreateGame(ctx context.Context, id int64, name string) error
	CreateGameInfo(ctx context.Context, info GameInfo) error
	FindUserGames(ctx context.Context, userId int64) ([]Game, error)
	WithTx(tx *sql.Tx) GamesRepository
}
