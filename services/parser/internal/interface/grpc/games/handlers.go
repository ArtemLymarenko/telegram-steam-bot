package gamesgrpc

import (
	"context"
	"github.com/ArtemLymarenko/steam-tg-bot/protos/gen/go/games"
	"github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GamesService interface {
	FindUserGames(ctx context.Context, userId int64) ([]domain.Game, error)
	AddUserGame(ctx context.Context, userId, gameId int64) error
	SearchGamesByName(ctx context.Context, name string) ([]domain.Game, error)
}

type serverApi struct {
	games.UnimplementedGamesServer
	gamesService GamesService
}

func NewServerApi(gamesService GamesService) *serverApi {
	return &serverApi{
		gamesService: gamesService,
	}
}

func (s *serverApi) Register(gRPC *grpc.Server) {
	games.RegisterGamesServer(gRPC, s)
}

func (s *serverApi) GetUserGames(
	ctx context.Context,
	req *games.GetUserGamesRequest,
) (*games.GetUserGamesResponse, error) {
	userGames, err := s.gamesService.FindUserGames(ctx, req.UserId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user games were not found")
	}

	getUserGamesResponse := &games.GetUserGamesResponse{}
	getUserGamesResponse.Games = make([]*games.Game, len(userGames))
	for i, game := range userGames {
		getUserGamesResponse.Games[i] = mapDomainGameToGameResponse(game)
	}

	return getUserGamesResponse, nil
}

func (s *serverApi) AddUserGame(
	ctx context.Context,
	req *games.AddUserGameRequest,
) (*games.AddUserGameResponse, error) {
	if req.UserId < 0 || req.GameId < 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	err := s.gamesService.AddUserGame(ctx, req.UserId, req.GameId)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, "game already exists")
	}

	return &games.AddUserGameResponse{
		Success: true,
	}, nil
}

func (s *serverApi) SearchGamesByName(
	ctx context.Context,
	req *games.SearchGamesByNameRequest,
) (*games.SearchGamesByNameResponse, error) {
	if req.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	foundGames, err := s.gamesService.SearchGamesByName(ctx, req.GetName())
	if err != nil {
		return nil, status.Error(codes.NotFound, "games were not found")
	}

	getUserGamesResponse := &games.SearchGamesByNameResponse{}
	getUserGamesResponse.Games = make([]*games.Game, len(foundGames))
	for i, game := range foundGames {
		getUserGamesResponse.Games[i] = mapDomainGameToGameResponse(game)
	}

	return getUserGamesResponse, nil
}
