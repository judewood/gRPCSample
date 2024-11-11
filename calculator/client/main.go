package main

import (
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/judewood/gRPCSample/calculator/proto"
)

var endpoint = "localhost:7777"

func main() {
	// create a persistent connection to a server endpoint
	conn, err := grpc.NewClient(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create connection. Error: %v", err)
	}
	//ensure the connection is closed before its enclosing function returns 
	defer conn.Close()

	// create an concrete  client struct from the generated code 
	c := pb.NewCalculatorServiceClient(conn)

	// Now call the server endpoints
	Sum(c, 1, 2)
	ops := []int64{1,2,3}
	SumMany(c,ops)
	CumulativeSum(c,ops)
	CountDown(c, 6)
	// Inputting a negative number  will return status InvalidArgument status code and the server supplied error description
	SquareRoot(c,13)
	// the server will respond after 3 seconds Deadline cna be adjusted to show a timeout 
	SumDelay(c, 10,20, 5 * time.Second)
}
