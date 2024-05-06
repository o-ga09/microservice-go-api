package grpc

import (
	"context"
	"os"

	authpb "github.com/o-ga09/microservice-go-api/services/auth/proto"
	entrancepb "github.com/o-ga09/microservice-go-api/services/entrance/proto"
	healthcheckpb "github.com/o-ga09/microservice-go-api/services/healthcheck/proto"
	paymentpb "github.com/o-ga09/microservice-go-api/services/payment/proto"
	userpb "github.com/o-ga09/microservice-go-api/services/user/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	}

	host := os.Getenv("AUTH_HOST")

	AuthCon, err := grpc.DialContext(ctx, host, opts...)
	if err != nil {
		return nil, err
	}

	authClient := authpb.NewAuthServiceClient(AuthCon)
	res, err := authClient.Health(ctx, &authpb.HealthCheckRequest{})
	if err != nil {
		return nil, err
	}

	return &healthcheckpb.HealthCheckResponse{
		Message: res.GetMessage(),
	}, nil
}

func (s *server) HealthCheckEntrance(ctx context.Context, req *healthcheckpb.HealthCheckRequest) (*healthcheckpb.HealthCheckResponse, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	}

	host := os.Getenv("ENTRANCE_HOST")

	EntranceCon, err := grpc.DialContext(ctx, host, opts...)
	if err != nil {
		return nil, err
	}

	EntranceClient := entrancepb.NewEntranceServiceClient(EntranceCon)
	res, err := EntranceClient.Health(ctx, &entrancepb.HealthCheckRequest{})
	if err != nil {
		return nil, err
	}

	return &healthcheckpb.HealthCheckResponse{
		Message: res.GetMessage(),
	}, nil
}

func (s *server) HealthCheckPayment(ctx context.Context, req *healthcheckpb.HealthCheckRequest) (*healthcheckpb.HealthCheckResponse, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	}

	host := os.Getenv("PAYMENT_HOST")

	PaymentCon, err := grpc.DialContext(ctx, host, opts...)
	if err != nil {
		return nil, err
	}

	PaymentClient := paymentpb.NewPaymentServiceClient(PaymentCon)
	res, err := PaymentClient.Health(ctx, &paymentpb.HealthCheckRequest{})
	if err != nil {
		return nil, err
	}

	return &healthcheckpb.HealthCheckResponse{
		Message: res.GetMessage(),
	}, nil
}

func (s *server) HealthCheckUser(ctx context.Context, req *healthcheckpb.HealthCheckRequest) (*healthcheckpb.HealthCheckResponse, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	}

	host := os.Getenv("USER_HOST")

	UserCon, err := grpc.DialContext(ctx, host, opts...)
	if err != nil {
		return nil, err
	}

	userClient := userpb.NewUserServiceClient(UserCon)
	res, err := userClient.Health(ctx, &userpb.HealthCheckRequest{})
	if err != nil {
		return nil, err
	}

	return &healthcheckpb.HealthCheckResponse{
		Message: res.GetMessage(),
	}, nil
}
