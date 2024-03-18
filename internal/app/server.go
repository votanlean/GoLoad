package app

import (
	"context"
	"log"

	"github.com/votanlean/GoLoad/internal/handler/grpc"
	"github.com/votanlean/GoLoad/internal/handler/http"
)

type Server struct {
	grpcServer grpc.Server
	httpServer http.Server
}

func NewServer(
	grpcServer grpc.Server,
	httpServer http.Server,
) *Server {
	return &Server{
		grpcServer: grpcServer,
		httpServer: httpServer,
	}
}

func (s *Server) Start() error {
	go func ()  {
		if err := s.grpcServer.Start(context.Background()); err != nil {
			log.Panic(err)
		}
	}()

	go func ()  {
		if err := s.httpServer.Start(context.Background()); err != nil {
			log.Panic(err)
		}
	}()

	return nil
}
