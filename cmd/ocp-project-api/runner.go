package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/lib/pq"
	projectApi "github.com/ozoncp/ocp-project-api/internal/api/ocp-project-api"
	"github.com/ozoncp/ocp-project-api/internal/producer"
	"github.com/ozoncp/ocp-project-api/internal/prom"
	"github.com/ozoncp/ocp-project-api/internal/storage"
	desc "github.com/ozoncp/ocp-project-api/pkg/ocp-project-api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort  = ":8082"
	httpPort  = ":8080"
	promPort  = ":9100"
	chunkSize = 1
)

func runGrpcAndGateway() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	db, err := storage.OpenDB()
	if err != nil {
		return fmt.Errorf("connect to db error: %v", err)
	}

	projectStorage := storage.NewProjectStorage(db, chunkSize)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	var logProducer producer.Producer
	logProducer, err = producer.NewProducer(ctx)
	if err != nil {
		return fmt.Errorf("Kafka producer creation failed: %v", err)
	}

	desc.RegisterOcpProjectApiServer(grpcServer, projectApi.NewOcpProjectApi(projectStorage, logProducer))
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return fmt.Errorf("initialization grpc server error: %v", err)
	}

	var group errgroup.Group
	group.Go(func() error {
		log.Info().Msgf("Serving grpc requests on %s", grpcPort)
		return grpcServer.Serve(listen)
	})

	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	group.Go(func() error {
		if err := desc.RegisterOcpProjectApiHandlerFromEndpoint(ctx, gwmux, grpcPort, opts); err != nil {
			return fmt.Errorf("registering gateway failed: %v", err)
		}

		mux := http.NewServeMux()
		mux.Handle("/", gwmux)

		log.Info().Msgf("Grpc gateway server listening on %s", httpPort)
		if err = http.ListenAndServe(httpPort, mux); err != nil {
			return fmt.Errorf("gateway http server failed: %v", err)
		}

		return nil
	})

	group.Go(func() error {
		prom.RegisterProjectMetrics()

		http.Handle("/metrics", promhttp.Handler())
		log.Info().Msgf("Prom http listening on %s", promPort)
		if err = http.ListenAndServe(promPort, nil); err != nil {
			return fmt.Errorf("prom http server failed: %v", err)
		}
		return nil
	})

	return group.Wait()
}
