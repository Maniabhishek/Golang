package grpcserver

import (
	"log"
	"net"
	"os"

	"github.com/Maniabhishek/Golang/i.GoProjects/serverWithEcho/configs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcServer struct {
	grpcserver *grpc.Server
	port       string
}

func NewGrpcServer() (*GrpcServer, error) {
	port := os.Getenv(configs.GRPC_PORT)
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	return &GrpcServer{
		port:       port,
		grpcserver: grpcServer,
	}, nil
}

func (g *GrpcServer) StartServer() error {

	lis, err := net.Listen("tcp", ":"+g.port)
	if err != nil {
		return err
	}
	log.Printf("%s %s", "grpc server started and listening on ", g.port)
	if err := g.grpcserver.Serve(lis); err != nil {
		log.Panic(err)
		return err
	}
	return nil
}
