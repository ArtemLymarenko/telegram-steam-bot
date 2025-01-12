package grpcapp

import (
	"fmt"
	gamesgrpc "github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/interface/grpc/games"
	"github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
}

func New(port int) *App {
	grpcServer := grpc.NewServer()

	gamesService := service.NewGames(nil, nil)
	gamesApi := gamesgrpc.NewServerApi(gamesService)
	gamesApi.Register(grpcServer)

	return &App{
		gRPCServer: grpcServer,
		port:       port,
	}
}

func (a *App) MustStart() {
	listen, err := net.Listen("tcp", a.getAddr())
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("gRPC server started on %s\n", a.getAddr())
	err = a.gRPCServer.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) Stop() {
	a.gRPCServer.GracefulStop()
	log.Println("gRPC server stopped!")
}

func (a *App) getAddr() string {
	return fmt.Sprintf(":%d", a.port)
}
