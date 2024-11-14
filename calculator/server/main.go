package main

import (
	"log"
	"net"

	pb "github.com/judewood/gRPCSample/calculator/proto"
	"github.com/judewood/gRPCSample/internal/consts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type CalcServer struct {
	pb.CalculatorServiceServer
}

func main() {
	// tell the runtime what port to listen on and the transport protocol to use
	listener := getListener()

	s := grpc.NewServer(getServerOptions()...)

	// register s as  s as being the concrete implementation of
	//CalculatorServiceServer defined in the generated grpc code
	pb.RegisterCalculatorServiceServer(s, &CalcServer{})
	//reflection.Register(s) uncomment when using evans cli

	//start our server
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to create calculator server.Error %v", err)
	}
}

// getListener returns a listener that monitors calls to our url
func getListener() net.Listener {
	listener, err := net.Listen(consts.TCP, consts.ServerUrl)
	if err != nil {
		log.Fatalf("failed to create calculator listener")
	}
	log.Printf("Listening on %s", consts.ServerUrl)
	return listener
}

// getServerOptions gets the options to configure our http2 server with
// currently only enables/disables SSL
func getServerOptions() []grpc.ServerOption {
	opts := []grpc.ServerOption{}
	if consts.UseSSL {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("failed to get ssl credentials. Error %v", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}
	return opts
}
