package grpc

import (
	"context"
	"log/slog"
	"os"

	"github.com/o-ga09/microservice-go-api/pkg/mysql"
	entrancepb "github.com/o-ga09/microservice-go-api/services/entrance/proto"
)

type server struct {
	entrancepb.UnimplementedEntranceServiceServer
}

func (s *server) Health(ctx context.Context, in *entrancepb.HealthCheckRequest) (*entrancepb.HealthCheckResponse, error) {
	slog.Info("Health")

	db := mysql.New(ctx, os.Getenv("ENTRANCE_DB"))
	result := db.Select("select curdate()")

	if result.Error != nil {
		return &entrancepb.HealthCheckResponse{Message: "DB NOT CONNECTED!"}, nil
	}
	return &entrancepb.HealthCheckResponse{Message: "OK Healthy!"}, nil
}

func (s *server) CheckIn(ctx context.Context, in *entrancepb.CheckInRequest) (*entrancepb.CheckInResponse, error) {
	slog.Info("CheckIn")
	return &entrancepb.CheckInResponse{}, nil
}

func (s *server) CheckOut(ctx context.Context, in *entrancepb.CheckOutRequest) (*entrancepb.CheckOutResponse, error) {
	slog.Info("CheckOut")
	return &entrancepb.CheckOutResponse{}, nil
}
