package main

import (
	"context"
	pb "grpc/proto"
	"io"
	"log"
	"time"
)

func CallSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%s", res.Message)
}

func CallHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("Bidirectional Stream")
	stream, err := client.SayHelloBiDirectionalStreaming(context.Background())

	if err != nil {
		log.Fatal("Could not send names", err)
	}

	// waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while streamming  %v", err)
			}
			log.Println("Message", message)
		}
		// close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	// <-waitc
	log.Println("Bidirectional streaming finished")
}
