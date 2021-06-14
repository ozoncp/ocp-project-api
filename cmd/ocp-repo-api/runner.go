package main

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	repoApi "github.com/ozoncp/ocp-project-api/internal/api/ocp-repo-api"
	"github.com/ozoncp/ocp-project-api/internal/producer"
	"github.com/ozoncp/ocp-project-api/internal/prom"
	"github.com/ozoncp/ocp-project-api/internal/storage"
	desc "github.com/ozoncp/ocp-project-api/pkg/ocp-repo-api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = ":8083"
	httpPort = ":8081"
	promPort = ":9101"

	chunkSize = 1
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

	var logProducer producer.Producer
	logProducer, err = producer.NewProducer(ctx)
	if err != nil {
		log.Error().Msgf("Kafka producer creation failed: %v", err)
	}

	desc.RegisterOcpRepoApiServer(grpcServer, repoApi.NewOcpRepoApi(repoStorage, logProducer))
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

	group.Go(func() error {
		prom.RegisterRepoMetrics()

		http.Handle("/metrics", promhttp.Handler())
		log.Info().Msgf("Prom http listening on %s", promPort)
		if err = http.ListenAndServe(promPort, nil); err != nil {
			log.Error().Msgf("Prom http server fails: %v", err)
			return err
		}
		return nil
	})

	return group.Wait()
}
