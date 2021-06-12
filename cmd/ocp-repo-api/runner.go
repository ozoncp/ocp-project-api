package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	repoApi "github.com/ozoncp/ocp-project-api/internal/api/ocp-repo-api"
	"github.com/ozoncp/ocp-project-api/internal/storage"
	desc "github.com/ozoncp/ocp-project-api/pkg/ocp-repo-api"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
)

const (
	grpcPort  = ":8083"
	httpPort  = ":8080"
	chunkSize = 10
)

func runGrpcAndGateway() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	db, err := sqlx.Connect("postgres", "user=lobanov dbname=ocp sslmode=disable")
	if err != nil {
		return err
	}

	repoStorage := storage.NewRepoStorage(db, chunkSize)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	desc.RegisterOcpRepoApiServer(grpcServer, repoApi.NewOcpRepoApi(repoStorage))
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
		if err := desc.RegisterOcpRepoApiHandlerFromEndpoint(ctx, gwmux, grpcPort, opts); err != nil {
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
