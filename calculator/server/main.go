package main

import (
	"log"
	"net"

	pb "github.com/judewood/gRPCSample/calculator/proto"
	"google.golang.org/grpc"
)

var endpoint = "0.0.0.0:7777"

type CalcServer struct {
	pb.CalculatorServiceServer
}

func main() {
	protocol := "tcp"
	// tell the runtime what port to listen on and the transport protocol to use 
	listener, err := net.Listen(protocol, endpoint)
	if err != nil {
		log.Fatalf("failed to create calculator listener")
	}
	log.Printf("Listening on %s with protocol %s", endpoint, protocol)

	s := grpc.NewServer()

	// register s as  s as being the concrete implementation of 
	//CalculatorServiceServer defined in the generated gprc code 
	pb.RegisterCalculatorServiceServer(s, &CalcServer{})

	//start our server and listen 
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to create sum server.Error %v", err)
	}
}
