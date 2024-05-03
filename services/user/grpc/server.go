package grpc

import (
	"context"
	"log/slog"

	gen "github.com/o-ga09/microservice-go-api/services/user/proto"
)

type server struct {
	gen.UnimplementedUserServiceServer
}

func (s *server) GetUsers(ctx context.Context, in *gen.UserRequest) (*gen.UserResponse, error) {
	slog.Info("GetUsers")
	return &gen.UserResponse{}, nil
}

func (s *server) GetUserById(ctx context.Context, in *gen.UserRequest) (*gen.UserResponse, error) {
	slog.Info("GetUserById")
	return &gen.UserResponse{}, nil
}

func (s *server) SaveUser(ctx context.Context, in *gen.UserRequest) (*gen.UserResponse, error) {
	slog.Info("SaveUser")
	return &gen.UserResponse{}, nil
}
