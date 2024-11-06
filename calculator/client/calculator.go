package main

import (
	"context"
	"log"

	pb "github.com/judewood/gRPCSample/calculator/proto"
)

func Sum(c pb.SumServiceClient, op1, op2 int64) {
	log.Printf("requesting sum of %d and %d", op1, op2)
	resp, err := c.Sum(context.Background(), &pb.SumRequest{
		Op1: 1, Op2: 2,
	})
	if err != nil {
		log.Fatalf("failed to request Sum. Error: %v", err)
	}
	log.Printf("Result of %d + %d is %d", op1, op1, resp.Result)
}
