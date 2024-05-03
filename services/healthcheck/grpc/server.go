package grpc

import (
	"context"

	gen "github.com/o-ga09/microservice-go-api/services/healthcheck/proto"
)

type server struct {
	gen.UnimplementedHealthCheckServiceServer
}

func (s *server) HealthCheck(ctx context.Context, req *gen.HealthCheckRequest) (*gen.HealthCheckResponse, error) {
	return &gen.HealthCheckResponse{
		Message: "OK Healthy!",
	}, nil
}

func (s *server) HealthCheckDB(ctx context.Context, req *gen.HealthCheckRequest) (*gen.HealthCheckResponse, error) {
	return &gen.HealthCheckResponse{
		Message: "OK DB Healthy!",
	}, nil
}
