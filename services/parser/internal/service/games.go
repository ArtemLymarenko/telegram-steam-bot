package service

import (
	"context"
	"database/sql"
	"github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/domain"
	txmanager "github.com/ArtemLymarenko/steam-tg-bot/services/parser/pkg/tx_manager"
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

func (g *Games) FindGame(ctx context.Context, gameId int64) (*domain.Game, error) {
	return g.gamesRepo.FindGame(ctx, gameId)
}

func (g *Games) CreateGame(ctx context.Context, gameId int64, name string) error {
	return g.gamesRepo.CreateGame(ctx, gameId, name)
}

func (g *Games) CreateGameInfo(ctx context.Context, info domain.GameInfo) error {
	_, err := g.FindGame(ctx, info.GameId)
	if err != nil {
		return err
	}

	return g.gamesRepo.CreateGameInfo(ctx, info)
}

func (g *Games) CreateGameWithInfo(ctx context.Context, game domain.Game) error {
	options := &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	}

	transaction := func(ctx context.Context, tx *sql.Tx) error {
		withTx := g.gamesRepo.WithTx(tx)
		err := withTx.CreateGame(ctx, game.Id, game.Name)
		if err != nil {
			return err
		}

		err = withTx.CreateGameInfo(ctx, game.GameInfo)
		return err
	}

	return g.txManager.Run(ctx, options, transaction)
}

func (g *Games) AddUserGame(ctx context.Context, userId, gameId int64) error {
	_, err := g.FindGame(ctx, gameId)
	if err != nil {
		return err
	}

	return g.AddUserGame(ctx, userId, gameId)
}

func (g *Games) FindUserGames(ctx context.Context, userId int64) ([]domain.Game, error) {
	return g.gamesRepo.FindUserGames(ctx, userId)
}

func (g *Games) DeleteGameById(ctx context.Context, gameId int64) error {
	return g.gamesRepo.DeleteGameById(ctx, gameId)
}
