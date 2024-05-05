package grpc

import (
	"context"
	"log/slog"

	entrancepb "github.com/o-ga09/microservice-go-api/services/entrance/proto"
)

type server struct {
	entrancepb.UnimplementedEntranceServiceServer
}

func (s *server) CheckIn(ctx context.Context, in *entrancepb.CheckInRequest) (*entrancepb.CheckInResponse, error) {
	slog.Info("CheckIn")
	return &entrancepb.CheckInResponse{}, nil
}

func (s *server) CheckOut(ctx context.Context, in *entrancepb.CheckOutRequest) (*entrancepb.CheckOutResponse, error) {
	slog.Info("CheckOut")
	return &entrancepb.CheckOutResponse{}, nil
}
