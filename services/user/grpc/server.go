package grpc

import (
	"context"
	"log/slog"
	"os"

	"github.com/o-ga09/microservice-go-api/pkg/mysql"
	userpb "github.com/o-ga09/microservice-go-api/services/user/proto"
)

type server struct {
	userpb.UnimplementedUserServiceServer
}

func (s *server) Health(ctx context.Context, in *userpb.HealthCheckRequest) (*userpb.HealthCheckResponse, error) {
	slog.Info("Health")

	db := mysql.New(ctx, os.Getenv("USER_DB"))
	result := db.Select("select curdate()")

	if result.Error != nil {
		return &userpb.HealthCheckResponse{Message: "DB NOT CONNECTED!"}, nil
	}

	return &userpb.HealthCheckResponse{Message: "OK Healthy!"}, nil
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
