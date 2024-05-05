package grpc

import (
	"context"
	"log/slog"

	userpb "github.com/o-ga09/microservice-go-api/services/user/proto"
)

type server struct {
	userpb.UnimplementedUserServiceServer
}

func (s *server) GetUsers(ctx context.Context, in *userpb.UserRequest) (*userpb.UserResponse, error) {
	slog.Info("GetUsers")
	return &userpb.UserResponse{}, nil
}

func (s *server) GetUserById(ctx context.Context, in *userpb.UserRequest) (*userpb.UserResponse, error) {
	slog.Info("GetUserById")
	return &userpb.UserResponse{}, nil
}

func (s *server) SaveUser(ctx context.Context, in *userpb.UserRequest) (*userpb.UserResponse, error) {
	slog.Info("SaveUser")
	return &userpb.UserResponse{}, nil
}
