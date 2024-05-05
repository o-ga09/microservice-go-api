package grpc

import (
	"context"
	"log/slog"

	paymentpb "github.com/o-ga09/microservice-go-api/services/payment/proto"
)

type server struct {
	paymentpb.UnimplementedPaymentServiceServer
}

func (s *server) RecordTransaction(ctx context.Context, in *paymentpb.RecordTransactionRequest) (*paymentpb.RecordTransactionResponse, error) {
	slog.Info("RecordTransaction")
	return &paymentpb.RecordTransactionResponse{}, nil
}

func (s *server) ProcessPayment(ctx context.Context, in *paymentpb.ProcessPaymentRequest) (*paymentpb.ProcessPaymentResponse, error) {
	slog.Info("ProcessPayment")
	return &paymentpb.ProcessPaymentResponse{}, nil
}
