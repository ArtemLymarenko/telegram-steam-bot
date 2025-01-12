package grpcapp

import (
	"fmt"
	gamesgrpc "github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/interface/grpc/games"
	"google.golang.org/grpc"
	"log"
	"net"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
}

func New(port int, gamesApi *gamesgrpc.ServerApi) *App {
	grpcServer := grpc.NewServer()

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
	log.Println("gRPC server is stopping...")
	a.gRPCServer.GracefulStop()
	log.Println("gRPC server stopped!")
}

func (a *App) getAddr() string {
	return fmt.Sprintf(":%d", a.port)
}
