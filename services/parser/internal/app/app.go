package app

import (
	grpcapp "github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/app/grpc"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	GRPCApp *grpcapp.App
}

func New(grpcPort int) *App {
	grpcApp := grpcapp.New(grpcPort)

	return &App{
		GRPCApp: grpcApp,
	}
}

func (a *App) Start() {
	stopCh := a.waitForShutdown()
	a.GRPCApp.MustStart()
	<-stopCh
}

func (a *App) gracefulStopLogic() {
	a.GRPCApp.Stop()
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
