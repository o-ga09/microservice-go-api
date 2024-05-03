package grpc

import (
	"context"
	"log/slog"

	gen "github.com/o-ga09/microservice-go-api/services/entrance/proto"
)

type server struct {
	gen.UnimplementedEntranceServiceServer
}

func (s *server) CheckIn(ctx context.Context, in *gen.CheckInRequest) (*gen.CheckInResponse, error) {
	slog.Info("CheckIn")
	return &gen.CheckInResponse{}, nil
}

func (s *server) CheckOut(ctx context.Context, in *gen.CheckOutRequest) (*gen.CheckOutResponse, error) {
	slog.Info("CheckOut")
	return &gen.CheckOutResponse{}, nil
}
