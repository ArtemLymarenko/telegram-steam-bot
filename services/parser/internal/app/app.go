package app

import (
	grpcapp "github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/app/grpc"
	repository "github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/infrastructure/sqlite"
	gamesgrpc "github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/interface/grpc/games"
	"github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/service/games"
	sqlite3Wrapper "github.com/ArtemLymarenko/steam-tg-bot/services/parser/pkg/sqlite3_wrapper"
	txmanager "github.com/ArtemLymarenko/steam-tg-bot/services/parser/pkg/tx_manager"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	gRPCApp *grpcapp.App
	sqlite  *sqlite3Wrapper.Sqlite3
}

func New(grpcPort int, sqlite *sqlite3Wrapper.Sqlite3) *App {
	db := sqlite.GetDbInstance()

	manager := txmanager.New(db)

	gamesRepo := repository.NewGames(db)

	gamesSvc := games_service.New(gamesRepo, manager)

	gamesGrpcApi := gamesgrpc.NewServerApi(gamesSvc)

	grpcApp := grpcapp.New(grpcPort, gamesGrpcApi)

	return &App{
		gRPCApp: grpcApp,
		sqlite:  sqlite,
	}
}

func (a *App) Start() {
	stopCh := a.waitForShutdown()
	a.gRPCApp.MustStart()
	<-stopCh
}

func (a *App) gracefulStopLogic() {
	a.sqlite.CloseConnection()
	a.gRPCApp.Stop()
}

func (a *App) waitForShutdown() <-chan struct{} {
	stopCh := make(chan struct{})
	go func() {
		signalCh := make(chan os.Signal, 1)
		signal.Notify(signalCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-signalCh
		a.gracefulStopLogic()
		close(stopCh)
	}()
	return stopCh
}
