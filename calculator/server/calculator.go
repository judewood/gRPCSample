package main

import (
	"context"
	"log"

	pb "github.com/judewood/gRPCSample/calculator/proto"
)

func (s *CalcServer) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {

	result := in.Op1 + in.Op2
	log.Printf("Received request %v.\n", in)
	resp := pb.SumResponse{
		Result: result,
	}
	log.Printf("Responding with %d.\n", resp.Result)
	return &resp, nil
}
