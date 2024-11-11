package main

import (
	"log"
	"net"

	pb "github.com/judewood/gRPCSample/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	opts := []grpc.ServerOption{}
	tls := true //true to use SSL  - must match client setting
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("failed to get ssl credentials. Error %v", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}
	s := grpc.NewServer(opts...)

	// register s as  s as being the concrete implementation of
	//CalculatorServiceServer defined in the generated grpc code
	pb.RegisterCalculatorServiceServer(s, &CalcServer{})

	//start our server and listen
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to create sum server.Error %v", err)
	}
}
