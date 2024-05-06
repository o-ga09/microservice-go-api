package grpc

import (
	"context"

	pkggrpc "github.com/o-ga09/microservice-go-api/pkg/grpc"
	paymentpb "github.com/o-ga09/microservice-go-api/services/payment/proto"
	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, port int) error {
	// oprs := []grpc.DialOption{
	// 	grpc.WithInsecure(),
	// 	grpc.WithBlock(),
	// 	grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	// }

	// con, err := grpc.DialContext(ctx, "localhost:"+port, oprs...)
	// if err != nil {
	// 	return err
	// }

	svc := &server{}

	return pkggrpc.NewServer(port, func(server *grpc.Server) {
		paymentpb.RegisterPaymentServiceServer(server, svc)
	}).Start(ctx)
}
