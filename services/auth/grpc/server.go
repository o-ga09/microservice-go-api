package grpc

import (
	"context"
	"log/slog"

	authpb "github.com/o-ga09/microservice-go-api/services/auth/proto"
)

type server struct {
	authpb.UnimplementedAuthServiceServer
}

func (s *server) SignUp(ctx context.Context, in *authpb.SignUpRequest) (*authpb.SignUpResponse, error) {
	slog.Info("SignUp")
	return &authpb.SignUpResponse{}, nil
}

func (s *server) SignIn(ctx context.Context, in *authpb.SignInRequest) (*authpb.SignInResponse, error) {
	slog.Info("SignIn")
	return &authpb.SignInResponse{}, nil
}

func (s *server) SignOut(ctx context.Context, in *authpb.SignOutRequest) (*authpb.SignOutResponse, error) {
	slog.Info("SignOut")
	return &authpb.SignOutResponse{}, nil
}
