package grpc

import (
	"context"
	"log/slog"

	gen "github.com/o-ga09/microservice-go-api/services/auth/proto"
)

type server struct {
	gen.UnimplementedAuthServiceServer
}

func (s *server) SignUp(ctx context.Context, in *gen.SignUpRequest) (*gen.SignUpResponse, error) {
	slog.Info("SignUp")
	return &gen.SignUpResponse{}, nil
}

func (s *server) SignIn(ctx context.Context, in *gen.SignInRequest) (*gen.SignInResponse, error) {
	slog.Info("SignIn")
	return &gen.SignInResponse{}, nil
}

func (s *server) SignOut(ctx context.Context, in *gen.SignOutRequest) (*gen.SignOutResponse, error) {
	slog.Info("SignOut")
	return &gen.SignOutResponse{}, nil
}
