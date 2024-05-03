package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/o-ga09/microservice-go-api/services/gateway/http"
)

func main() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	httpErrCh := make(chan error, 1)

	go func() {
		httpErrCh <- http.RunServer(ctx, 8081)
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
