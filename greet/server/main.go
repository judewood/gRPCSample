package main

import (
	"log"
	"net"

	pb "github.com/judewood/gRPCSample/greet/proto"
	"google.golang.org/grpc"
)

var endpoint string = "0.0.0.0:6666" //localhost and port

type Server struct {
	pb.GreetServiceServer
}

func main() {
	listener, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to create listener for endpoint %s. Error: %v\n", endpoint, err)
	}
	log.Printf("Listening on port %s... ", endpoint)
	s := grpc.NewServer()

	pb.RegisterGreetServiceServer(s, &Server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to create server.Error %v", err)
	}
}
