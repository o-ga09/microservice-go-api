package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	port := os.Getenv("AUTHSERVICE_PORT")
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}
	defer listen.Close()

	s := grpc.NewServer()

	go func() {
		log.Printf("sterting server on port %s", port)
		s.Serve(listen)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Printf("shutting down server ...")
	s.GracefulStop()

	return nil
}
