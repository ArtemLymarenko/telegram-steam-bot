package gamesgrpc

import (
	"context"
	"errors"
	"github.com/ArtemLymarenko/steam-tg-bot/protos/gen/go/games"
	"github.com/ArtemLymarenko/steam-tg-bot/services/bot/internal/interface/dto"
	"github.com/ArtemLymarenko/steam-tg-bot/services/bot/internal/interface/mapper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type ClientApi struct {
	games.GamesClient
}

func NewClientApi(gRPC *grpc.ClientConn) *ClientApi {
	client := games.NewGamesClient(gRPC)
	return &ClientApi{
		GamesClient: client,
	}
}

func (c *ClientApi) GetUserGames(
	ctx context.Context,
	req dto.GetUserGamesRequest,
) (*dto.GetUserGamesResponse, error) {
	userGames, err := c.GamesClient.GetUserGames(ctx, &games.GetUserGamesRequest{UserId: req.UserId})
	if err != nil {
		s, _ := status.FromError(err)
		return nil, errors.New(s.Message())
	}

	mappedGames := mapper.GamesGrpcToDto(userGames.Games)
	return &dto.GetUserGamesResponse{Games: mappedGames}, nil
}

func (c *ClientApi) AddUserGame(
	ctx context.Context,
	req dto.AddUserGameRequest,
) (dto.AddUserGameResponse, error) {
	result, err := c.GamesClient.AddUserGame(ctx, &games.AddUserGameRequest{
		UserId: req.UserId,
		GameId: req.GameId,
	})
	if err != nil {
		s, _ := status.FromError(err)
		return dto.AddUserGameResponse{Success: false}, errors.New(s.Message())
	}

	if result == nil {
		return dto.AddUserGameResponse{Success: false}, errors.New("failed to add game")
	}

	return dto.AddUserGameResponse{Success: result.Success}, nil
}

func (c *ClientApi) DeleteUserGame(
	ctx context.Context,
	req dto.DeleteUserGameRequest,
) (dto.DeleteUserGameResponse, error) {
	result, err := c.GamesClient.DeleteUserGame(ctx, &games.DeleteUserGameRequest{
		UserId: req.UserId,
		GameId: req.GameId,
	})
	if err != nil {
		s, _ := status.FromError(err)
		return dto.DeleteUserGameResponse{Success: false}, errors.New(s.Message())
	}
	if result == nil {
		return dto.DeleteUserGameResponse{Success: false}, errors.New("failed to add game")
	}

	return dto.DeleteUserGameResponse{Success: result.Success}, nil
}

func (c *ClientApi) SearchGamesByName(
	ctx context.Context,
	req dto.SearchGameRequest,
) (*dto.SearchGameResponse, error) {
	result, err := c.GamesClient.SearchGamesByName(ctx, &games.SearchGamesByNameRequest{
		Name: req.Name,
	})
	if err != nil {
		s, _ := status.FromError(err)
		return nil, errors.New(s.Message())
	}

	mappedGames := mapper.GamesGrpcToDto(result.Games)
	return &dto.SearchGameResponse{Games: mappedGames}, nil
}
