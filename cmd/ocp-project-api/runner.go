package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/ozoncp/ocp-project-api/internal/api"
	desc "github.com/ozoncp/ocp-project-api/pkg/ocp-project-api"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

const (
	grpcPort = ":8082"
	httpPort = ":8080"
)

func runGrpcAndGateway() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcServer := grpc.NewServer()
	desc.RegisterOcpProjectApiServer(grpcServer, api.NewOcpProjectApi())
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Error().Msgf("Grpc server error: %v", err)
		return err
	}

	var group errgroup.Group
	group.Go(func() error {
		log.Info().Msg("Serving grpc requests...")
		return grpcServer.Serve(listen)
	})

	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	group.Go(func() error {
		if err := desc.RegisterOcpProjectApiHandlerFromEndpoint(ctx, gwmux, grpcPort, opts); err != nil {
			log.Error().Msgf("Register gateway fails: %v", err)
			return err
		}

		mux := http.NewServeMux()
		mux.Handle("/", gwmux)

		log.Info().Msgf("Http server listening on %s", httpPort)
		if err = http.ListenAndServe(httpPort, mux); err != nil {
			log.Error().Msgf("Gateway http server fails: %v", err)
			return err
		}

		return nil
	})

	return group.Wait()
}
