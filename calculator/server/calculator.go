package main

import (
	"context"
	"io"
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

func (s *CalcServer) SumMany(streamIn grpc.ClientStreamingServer[pb.SumManyRequest, pb.SumManyResponse]) error {
	var sum int64 = 0
	for {
		req, err := streamIn.Recv()
		if err == io.EOF {
			resp := pb.SumManyResponse{
				Op1: sum,
			}
			return streamIn.SendAndClose(&resp)
		}
		if err != nil {
			log.Fatalf("failed to read stream request for SumMany %v", err)
		}
		log.Printf("Received request %v.\n", req.Op1)
		sum += req.Op1
	}
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
	for i := n - 1; i >= 0; i-- {
		values[index] = i
		index++
	}
	log.Printf("values %#v", values)
	return values
}
