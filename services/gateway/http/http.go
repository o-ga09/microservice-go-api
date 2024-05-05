package http

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	Authgen "github.com/o-ga09/microservice-go-api/services/auth/proto"
	Entrancegen "github.com/o-ga09/microservice-go-api/services/entrance/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func RunServer(ctx context.Context, port int) error {
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:   true,
					EmitUnpopulated: true,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		}),
	)

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	}
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	AuthCon, err := grpc.DialContext(ctx, "authsrv01:8081", opts...)
	if err != nil {
		return err
	}

	EntranceCon, err := grpc.DialContext(ctx, "entrancesrv01:8083", opts...)
	if err != nil {
		return err
	}

	if err := Authgen.RegisterAuthServiceHandlerClient(ctx, mux, Authgen.NewAuthServiceClient(AuthCon)); err != nil {
		return err
	}

	if err := Entrancegen.RegisterEntranceServiceHandlerClient(ctx, mux, Entrancegen.NewEntranceServiceClient(EntranceCon)); err != nil {
		return err
	}

	errCh := make(chan error, 1)

	go func() {
		slog.Info(fmt.Sprintf("Starting gRPC server on port %d", port))
		errCh <- server.ListenAndServe()
	}()

	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
		return nil
	case <-ctx.Done():
		if err := server.Shutdown(ctx); err != nil {
			return err
		}
		if err := <-errCh; err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	}
}
