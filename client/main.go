package main

import (
	pb "grpc/proto"
	"log"

	"google.golang.org/grpc"
)

const (
	port = ":8000"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("unable to connect %v", err)
	}

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Akash", "Akash1", "Akash2"},
	}

	// CallSayHello(client)
	CallHelloBidirectionalStream(client, names)

	defer conn.Close()
}
