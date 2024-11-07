package main

import (
	"context"
	"io"
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

func DoGreetMany(c pb.GreetServiceClient) {
	log.Println("doGreetMany was invoked")
	req :=  pb.GreetRequest{
		FirstName: "Freds",
	}
	stream, err := c.GreetMany(context.Background(), &req)
	if err != nil {
		log.Fatalf("failed to send greet many request. Error %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			log.Println("End of stream")
			break
		}
		if err != nil {
			log.Fatalf("failed during stream. Error %v", err)
		}
		log.Printf("response received %s", msg.Result)
	}
}
