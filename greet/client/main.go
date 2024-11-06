package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/judewood/gRPCSample/greet/proto"
)

var endpoint string = "localhost:6666" //localhost not 0.0.0.0 because this is the client

func main() {
	// gRPC uses ssl be default - we are bypassing this for now with insecure.NewCredentials
	conn, err := grpc.NewClient(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to  endpoint. Error: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	DoGreet(c)
}
