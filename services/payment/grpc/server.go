package grpc

import (
	"context"
	"log/slog"
	"os"

	"github.com/o-ga09/microservice-go-api/pkg/mysql"
	paymentpb "github.com/o-ga09/microservice-go-api/services/payment/proto"
)

type server struct {
	paymentpb.UnimplementedPaymentServiceServer
}

func (s *server) Health(ctx context.Context, in *paymentpb.HealthCheckRequest) (*paymentpb.HealthCheckResponse, error) {
	slog.Info("Health")

	db := mysql.New(ctx, os.Getenv("PAYMENT_DB"))
	result := db.Select("select curdate()")

	if result.Error != nil {
		return &paymentpb.HealthCheckResponse{Message: "DB NOT CONNECTED!"}, nil
	}

	return &paymentpb.HealthCheckResponse{Message: "OK Healthy!"}, nil
}

func (s *server) RecordTransaction(ctx context.Context, in *paymentpb.RecordTransactionRequest) (*paymentpb.RecordTransactionResponse, error) {
	slog.Info("RecordTransaction")
	return &paymentpb.RecordTransactionResponse{}, nil
}

func (s *server) ProcessPayment(ctx context.Context, in *paymentpb.ProcessPaymentRequest) (*paymentpb.ProcessPaymentResponse, error) {
	slog.Info("ProcessPayment")
	return &paymentpb.ProcessPaymentResponse{}, nil
}
