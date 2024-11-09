package main

import (
	"context"
	"io"
	"log"

	pb "github.com/judewood/gRPCSample/calculator/proto"
)

func Sum(c pb.CalculatorServiceClient, op1, op2 int64) {
	log.Printf("requesting sum of %d and %d", op1, op2)
	resp, err := c.Sum(context.Background(), &pb.SumRequest{
		Op1: 1, Op2: 2,
	})
	if err != nil {
		log.Fatalf("failed to request Sum. Error: %v", err)
	}
	log.Printf("Result of %d + %d is %d", op1, op2, resp.Result)
}


func CountDown(c pb.CalculatorServiceClient, val int64) {
	log.Printf("requesting Countdown from %d", val)
	req := pb.CountDownRequest{Value: val}
	stream, err := c.CountDown(context.Background(), &req)
	if err != nil {
		log.Fatalf("failed to send countdown  request. Error %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			log.Printf("end of stream")
			break
		}
		if err != nil {
			log.Fatalf("failed during stream. Error %v", err)
		}
		log.Printf("received:  %d", msg.Count)
	}
}
