package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/ozoncp/ocp-project-api/internal/api"
	desc "github.com/ozoncp/ocp-project-api/pkg/ocp-project-api"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

const (
	grpcPort = ":8082"
	httpPort = ":8080"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcServer := grpc.NewServer()
	desc.RegisterOcpProjectApiServer(grpcServer, api.NewOcpProjectApi())
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var group errgroup.Group
	group.Go(func() error {
		fmt.Println("Serving grpc requests...")
		return grpcServer.Serve(listen)
	})

	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	group.Go(func() error {
		if err := desc.RegisterOcpProjectApiHandlerFromEndpoint(ctx, gwmux, grpcPort, opts); err != nil {
			fmt.Printf("Register gateway fails: %v\n", err)
			return err
		}

		mux := http.NewServeMux()
		mux.Handle("/", gwmux)

		fmt.Printf("Http server listening on %s\n", httpPort)
		if err = http.ListenAndServe(httpPort, mux); err != nil {
			fmt.Printf("Gateway http server fails: %v\n", err)
			return err
		}

		return nil
	})

	return group.Wait()
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		return
	}
}
