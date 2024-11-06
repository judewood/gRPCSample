package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/judewood/gRPCSample/calculator/proto"
)

var endpoint = "localhost:7777"

func main() {
	conn, err := grpc.NewClient(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create connection. Error: %v", err)
	}
	defer conn.Close()

	c := pb.NewSumServiceClient(conn)

	Sum(c, 1, 2)
}
