package grpc

import (
	"github.com/votanlean/GoLoad/internal/generated/grpc/go_load"
)

type Handler struct {
	go_load.UnimplementedGoLoadServiceServer
}

func NewHandler() go_load.GoLoadServiceServer {
	return &Handler{}
}