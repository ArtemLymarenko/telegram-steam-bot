package games_service

import (
	"context"
	"database/sql"
	"github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/domain/game"
	txmanager "github.com/ArtemLymarenko/steam-tg-bot/services/parser/pkg/tx_manager"
)

type Games struct {
	gamesRepo game.Repository
	txManager txmanager.TxManager
}

func New(repo game.Repository, tx txmanager.TxManager) *Games {
	return &Games{
		gamesRepo: repo,
		txManager: tx,
	}
}

func (g *Games) FindGame(ctx context.Context, gameId game.Id) (game.Game, error) {
	if err := gameId.Validate(); err != nil {
		return game.Game{}, err
	}

	return g.gamesRepo.FindGame(ctx, gameId)
}

func (g *Games) CreateGame(ctx context.Context, gameId game.Id, name game.Name) error {
	if err := gameId.Validate(); err != nil {
		return err
	}

	if err := name.Validate(); err != nil {
		return err
	}

	return g.gamesRepo.CreateGame(ctx, gameId, name)
}

func (g *Games) CreateGameInfo(ctx context.Context, info game.Info) error {
	if err := info.Validate(); err != nil {
		return err
	}

	_, err := g.FindGame(ctx, info.GameId)
	if err != nil {
		return err
	}

	return g.gamesRepo.CreateGameInfo(ctx, info)
}

func (g *Games) CreateGameWithInfo(ctx context.Context, game game.Game) error {
	if err := game.Validate(); err != nil {
		return err
	}

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

		err = withTx.CreateGameInfo(ctx, game.Info)
		return err
	}

	return g.txManager.Run(ctx, options, transaction)
}

func (g *Games) AddUserGame(ctx context.Context, userId game.UserId, gameId game.Id) error {
	if err := gameId.Validate(); err != nil {
		return err
	}

	_, err := g.FindGame(ctx, gameId)
	if err != nil {
		return err
	}

	return g.gamesRepo.AddUserGame(ctx, userId, gameId)
}

func (g *Games) DeleteUserGame(ctx context.Context, userId game.UserId, gameId game.Id) (game.Id, error) {
	if err := gameId.Validate(); err != nil {
		return 0, err
	}

	return g.gamesRepo.DeleteUserGame(ctx, userId, gameId)
}

func (g *Games) FindUserGames(ctx context.Context, userId game.UserId) ([]game.Game, error) {
	return g.gamesRepo.FindUserGames(ctx, userId)
}

func (g *Games) DeleteGameById(ctx context.Context, gameId game.Id) error {
	if err := gameId.Validate(); err != nil {
		return err
	}

	return g.gamesRepo.DeleteGameById(ctx, gameId)
}

func (g *Games) SearchGamesByName(ctx context.Context, name game.Name) ([]game.Game, error) {
	if err := name.Validate(); err != nil {
		return nil, err
	}

	found, err := g.gamesRepo.SearchGamesByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return found, nil
}
