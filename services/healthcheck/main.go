package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/o-ga09/microservice-go-api/services/healthcheck/grpc"
	"golang.org/x/sys/unix"
)

func main() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	ctx, stop := signal.NotifyContext(ctx, unix.SIGTERM, unix.SIGINT)
	defer stop()

	errCh := make(chan error, 1)
	go func() {
		errCh <- grpc.RunServer(ctx, 8080)
	}()

	select {
	case err := <-errCh:
		fmt.Println(err.Error())
		return 1
	case <-ctx.Done():
		fmt.Println("Shutting down...")
		return 0
	}
}
