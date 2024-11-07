package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/judewood/gRPCSample/greet/proto"
	"google.golang.org/grpc"
)

// Greet is the server side implementation of the Greet function
// defined in greet_grpc.pb.go
func (s *Server) Greet(ctx context.Context, request *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Received request %v.\n", request)
	resp := pb.GreetResponse{
		Result: "Hi there " + request.FirstName,
	}
	log.Printf("Responding with %s.\n", resp.Result)
	return &resp, nil
}

func (s *Server) GreetMany(in *pb.GreetRequest, stream grpc.ServerStreamingServer[pb.GreetResponse]) error {
	log.Printf("Received request %v.\n", in)

	for i := 0; i < 5; i++ {
		resp := pb.GreetResponse{
			Result: fmt.Sprintf("Hi there %s. Times %d", in.FirstName, i+1),
		}
		stream.Send(&resp)
	}
	return nil
}
