package main

import (
	"log"
	"net"

	pb "github.com/judewood/gRPCSample/calculator/proto"
	"google.golang.org/grpc"
)

var endpoint = "0.0.0.0:7777"

type CalcServer struct {
	pb.SumServiceServer
}

func main() {
	listener, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to create calculator listener")
	}
	log.Printf("Listening on %s", endpoint)

	s := grpc.NewServer()

	pb.RegisterSumServiceServer(s, &CalcServer{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to create sum server.Error %v", err)
	}
}
