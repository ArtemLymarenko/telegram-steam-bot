package gamesgrpc

import (
	"context"
	"errors"
	"github.com/ArtemLymarenko/steam-tg-bot/protos/gen/go/games"
	"github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/domain"
	"github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/domain/game"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GamesService interface {
	FindUserGames(ctx context.Context, userId game.UserId) ([]game.Game, error)
	AddUserGame(ctx context.Context, userId game.UserId, gameId game.Id) error
	DeleteUserGame(ctx context.Context, userId game.UserId, gameId game.Id) (game.Id, error)
	SearchGamesByName(ctx context.Context, name game.Name) ([]game.Game, error)
}

type ServerApi struct {
	games.UnimplementedGamesServer
	gamesService GamesService
}

func NewServerApi(gamesService GamesService) *ServerApi {
	return &ServerApi{
		gamesService: gamesService,
	}
}

func (s *ServerApi) Register(gRPC *grpc.Server) {
	games.RegisterGamesServer(gRPC, s)
}

func (s *ServerApi) GetUserGames(
	ctx context.Context,
	req *games.GetUserGamesRequest,
) (*games.GetUserGamesResponse, error) {
	userGames, err := s.gamesService.FindUserGames(ctx, game.UserId(req.UserId))
	if err != nil {
		var validationError domain.ValidationError
		if errors.As(err, &validationError) {
			return nil, status.Error(codes.InvalidArgument, validationError.Error())
		}

		return nil, status.Error(codes.NotFound, "user games were not found")
	}

	getUserGamesResponse := &games.GetUserGamesResponse{}
	getUserGamesResponse.Games = make([]*games.Game, len(userGames))
	for i, g := range userGames {
		getUserGamesResponse.Games[i] = mapDomainGameToGameResponse(g)
	}

	return getUserGamesResponse, nil
}

func (s *ServerApi) AddUserGame(
	ctx context.Context,
	req *games.AddUserGameRequest,
) (*games.AddUserGameResponse, error) {
	err := s.gamesService.AddUserGame(ctx, game.UserId(req.UserId), game.Id(req.GameId))
	if err != nil {
		var validationError domain.ValidationError
		if errors.As(err, &validationError) {
			return &games.AddUserGameResponse{
				Success: false,
			}, status.Error(codes.InvalidArgument, validationError.Error())
		}

		return &games.AddUserGameResponse{
			Success: false,
		}, status.Error(codes.AlreadyExists, "game already exists")
	}
	return &games.AddUserGameResponse{
		Success: true,
	}, nil
}

func (s *ServerApi) DeleteUserGame(
	ctx context.Context,
	req *games.DeleteUserGameRequest,
) (*games.DeleteUserGameResponse, error) {
	_, err := s.gamesService.DeleteUserGame(ctx, game.UserId(req.UserId), game.Id(req.GameId))
	if err != nil {
		var validationError domain.ValidationError
		if errors.As(err, &validationError) {
			return &games.DeleteUserGameResponse{
				Success: false,
			}, status.Error(codes.InvalidArgument, validationError.Error())
		}

		return &games.DeleteUserGameResponse{
			Success: false,
		}, status.Error(codes.Internal, "game does not	 exists")
	}
	return &games.DeleteUserGameResponse{
		Success: true,
	}, nil
}

func (s *ServerApi) SearchGamesByName(
	ctx context.Context,
	req *games.SearchGamesByNameRequest,
) (*games.SearchGamesByNameResponse, error) {
	foundGames, err := s.gamesService.SearchGamesByName(ctx, game.Name(req.GetName()))
	if err != nil {
		var validationError domain.ValidationError
		if errors.As(err, &validationError) {
			return nil, status.Error(codes.InvalidArgument, validationError.Error())
		}

		return nil, status.Error(codes.NotFound, "games were not found")
	}

	getUserGamesResponse := &games.SearchGamesByNameResponse{}
	getUserGamesResponse.Games = make([]*games.Game, len(foundGames))
	for i, g := range foundGames {
		getUserGamesResponse.Games[i] = mapDomainGameToGameResponse(g)
	}

	return getUserGamesResponse, nil
}
