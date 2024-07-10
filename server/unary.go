package main

import (
	"context"
	pb "grpc/proto"
	"io"
	"log"
)

func (s *HelloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}

func (s *HelloServer) SayHelloBiDirectionalStreaming(stream pb.GreetService_SayHelloBiDirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello" + " " + req.Name,
		}
		log.Println("Resp", res)
		if err := stream.Send(res); err != nil {
			log.Println("Error sending response", err)
			return err
		}
	}

}
