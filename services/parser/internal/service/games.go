package service

import (
	"context"
	"database/sql"
	"parser/internal/domain"
	txmanager "parser/pkg/tx_manager"
)

type Games struct {
	gamesRepo domain.GamesRepository
	txManager txmanager.TxManager
}

func NewGames(repo domain.GamesRepository, tx txmanager.TxManager) *Games {
	return &Games{
		gamesRepo: repo,
		txManager: tx,
	}
}

func (g *Games) TxTest() error {
	ctx := context.Background()
	return g.txManager.Run(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {
		txRepo := g.gamesRepo.WithTx(tx)
		err := txRepo.CreateGame(ctx, 2, "uncharted 2")
		if err != nil {
			return err
		}
		info := domain.GameInfo{
			GameId:          2,
			ImageUrl:        "",
			InitialPrice:    0,
			FinalPrice:      0,
			DiscountPercent: 0,
		}

		return txRepo.CreateGameInfo(ctx, info)
	})
}
