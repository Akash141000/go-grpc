package main

import (
	pb "grpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":8000"
)

type HelloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to start the server %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &HelloServer{})
	log.Printf("server started at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start the grpc server %v", err)
	}
}
