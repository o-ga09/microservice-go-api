package grpc

import (
	"context"

	healthcheckpb "github.com/o-ga09/microservice-go-api/services/healthcheck/proto"
)

type server struct {
	healthcheckpb.UnimplementedHealthCheckServiceServer
}

func (s *server) HealthCheck(ctx context.Context, req *healthcheckpb.HealthCheckRequest) (*healthcheckpb.HealthCheckResponse, error) {
	return &healthcheckpb.HealthCheckResponse{
		Message: "OK Healthy!",
	}, nil
}

func (s *server) HealthCheckDB(ctx context.Context, req *healthcheckpb.HealthCheckRequest) (*healthcheckpb.HealthCheckResponse, error) {
	return &healthcheckpb.HealthCheckResponse{
		Message: "OK DB Healthy!",
	}, nil
}

func (s *server) HealthCheckAuth(ctx context.Context, req *healthcheckpb.HealthCheckRequest) (*healthcheckpb.HealthCheckResponse, error) {
	return &healthcheckpb.HealthCheckResponse{
		Message: "OK Auth Healthy!",
	}, nil
}

func (s *server) HealthCheckEntrance(ctx context.Context, req *healthcheckpb.HealthCheckRequest) (*healthcheckpb.HealthCheckResponse, error) {
	return &healthcheckpb.HealthCheckResponse{
		Message: "OK Entrance Healthy!",
	}, nil
}

func (s *server) HealthCheckPayment(ctx context.Context, req *healthcheckpb.HealthCheckRequest) (*healthcheckpb.HealthCheckResponse, error) {
	return &healthcheckpb.HealthCheckResponse{
		Message: "OK Payment Healthy!",
	}, nil
}

func (s *server) HealthCheckUser(ctx context.Context, req *healthcheckpb.HealthCheckRequest) (*healthcheckpb.HealthCheckResponse, error) {
	return &healthcheckpb.HealthCheckResponse{
		Message: "OK User Healthy!",
	}, nil
}
