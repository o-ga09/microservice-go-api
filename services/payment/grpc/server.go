package grpc

import (
	"context"
	"log/slog"

	gen "github.com/o-ga09/microservice-go-api/services/payment/proto"
)

type server struct {
	gen.UnimplementedPaymentServiceServer
}

func (s *server) RecordTransaction(ctx context.Context, in *gen.RecordTransactionRequest) (*gen.RecordTransactionResponse, error) {
	slog.Info("RecordTransaction")
	return &gen.RecordTransactionResponse{}, nil
}

func (s *server) ProcessPayment(ctx context.Context, in *gen.ProcessPaymentRequest) (*gen.ProcessPaymentResponse, error) {
	slog.Info("ProcessPayment")
	return &gen.ProcessPaymentResponse{}, nil
}
