package grpc

import "google.golang.org/grpc"

type ServiceBindFunc[S any] func(s *grpc.Server, srv S)

type Server[S any] struct {
}

func NewServer[T any]() *Server[T] {
	return &Server[T]{}
}
