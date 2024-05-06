package main

import (
	"context"
	"log/slog"
	"os"
	"strconv"

	"github.com/o-ga09/microservice-go-api/services/payment/grpc"
)

func main() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	httpErrCh := make(chan error, 1)

	p := os.Getenv("PORT")
	port, _ := strconv.Atoi(p)

	go func() {
		httpErrCh <- grpc.RunServer(ctx, port)
	}()

	select {
	case err := <-httpErrCh:
		if err != nil {
			slog.InfoContext(ctx, "error", err)
			return 1
		}
	case <-ctx.Done():
		slog.InfoContext(ctx, "shutting down")
		return 0
	}

	return 0
}
