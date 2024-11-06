package main

import (
	"context"
	"log"

	pb "github.com/judewood/gRPCSample/greet/proto"
)

func DoGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Fred",
	})
	if err != nil {
		log.Fatalf("failed to send greet request. Error %v", err)
	}
	log.Printf("response received %s", res.GetResult())
}
