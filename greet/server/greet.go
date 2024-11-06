package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
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
