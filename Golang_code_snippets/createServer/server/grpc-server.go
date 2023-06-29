package server

import (
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	port   string
	Server *grpc.Server
}

func NewGrpcServer() (*grpcServer, error) {
	port := os.Getenv("GRPC_SERVER")

	server := grpc.NewServer()

	reflection.Register(server)

	return &grpcServer{
		port:   port,
		Server: server,
	}, nil
}

func (g *grpcServer) StartGrpcServer() error {
	listen, err := net.Listen("tcp", ":"+g.port)

	if err != nil {
		return nil
	}

	if err := g.Server.Serve(listen); err != nil {
		return nil
	}

	return nil
}
