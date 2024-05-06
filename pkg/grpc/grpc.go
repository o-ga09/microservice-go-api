package grpc

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	channelz "google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	server *grpc.Server
	port   int
}

func NewServer(port int, register func(*grpc.Server)) *Server {
	interceptors := []grpc.UnaryServerInterceptor{}

	opts := []grpc.ServerOption{
		middleware.WithUnaryServerChain(interceptors...),
	}

	server := grpc.NewServer(opts...)
	register(server)
	reflection.Register(server)
	channelz.RegisterChannelzServiceToServer(server)

	return &Server{
		server: server,
		port:   port,
	}
}

func (s *Server) Start(ctx context.Context) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}

	errCh := make(chan error)
	go func() {
		slog.Info(fmt.Sprintf("Starting gRPC server on port %d", s.port))
		if err := s.server.Serve(listener); err != nil {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
		return nil
	case <-ctx.Done():
		s.server.GracefulStop()
		return <-errCh
	}
}
