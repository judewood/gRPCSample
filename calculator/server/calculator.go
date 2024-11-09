package main

import (
	"context"
	"log"

	pb "github.com/judewood/gRPCSample/calculator/proto"
	"google.golang.org/grpc"
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

// CountDown returns all the values from counting down form initial value to zero
func (s *CalcServer) CountDown(in *pb.CountDownRequest, stream grpc.ServerStreamingServer[pb.CountDownResponse]) error {
	log.Printf("Received request %v.\n", in)

	values := countDown(in.Value)
	for _, v := range values {
		resp := pb.CountDownResponse{
			Count: v,
		}
		stream.Send(&resp)
	}
	return nil
}


func countDown(n int64) []int64 {
	values := make([]int64, n)
	index := 0
	for i := n-1; i >= 0 ; i-- {
           values[index] = i
		   index++
	}
	log.Printf("values %#v", values)
	return values
}
