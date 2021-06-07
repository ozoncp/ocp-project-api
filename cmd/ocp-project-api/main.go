package main

import (
	"fmt"
	"github.com/ozoncp/ocp-project-api/internal/api"
	desc "github.com/ozoncp/ocp-project-api/pkg/ocp-project-api"
	"google.golang.org/grpc"
	"net"
)

const (
	port = ":8080"
)

func main() {
	server := grpc.NewServer()
	desc.RegisterOcpProjectApiServer(server, api.NewOcpProjectApi())
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Serving requests...")
	server.Serve(listen)
}
