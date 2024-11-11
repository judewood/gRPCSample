package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"strconv"
	"time"

	pb "github.com/judewood/gRPCSample/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Sum receives a single (unary) request containing two int64s to be added together
// It returns a single response containing the result
// The signature of this method must match the equivalent in the generated CalculatorServiceServer interface so our server is implementing the handler method
func (s *CalcServer) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {

	result := in.Op1 + in.Op2
	log.Printf("Received request %v.\n", in)
	resp := pb.SumResponse{
		Result: result,
	}
	log.Printf("Responding with %d.\n", resp.Result)
	return &resp, nil
}

// SumDelay receives a single (unary) request containing two int64s to be added together
// It returns a single response containing the result after a delay (to show sample of deadline)
// The signature of this method must match the equivalent in the generated CalculatorServiceServer interface so our server is implementing the handler method
func (s *CalcServer) SumDelay(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Received SumDelay request %v .\n", in)
	delay := 3
	for i := 0; i < delay; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			reason := "Client cancelled the request"
			log.Println(reason)
			return nil, status.Error(codes.Canceled, reason)
		}
		log.Println("Sleeping for one second")
		time.Sleep(time.Second)
	}
	result := in.Op1 + in.Op2
	log.Println(result)
	resp := pb.SumResponse{
		Result: result,
	}
	log.Printf("SumDelay Responding with %d + %d  = %d.\n", in.Op1, in.Op2, resp.Result)
	return &resp, nil
}

// SumMany receives a stream of requests each containing a single int64 and creates a cumulative sum of these
// When the sender closes the stream (EOF) this method returns the sum
// The signature of this method must match the equivalent in the generated CalculatorServiceServer interface so our server is implementing the handler method
func (s *CalcServer) SumMany(inStream grpc.ClientStreamingServer[pb.SumManyRequest, pb.SumManyResponse]) error {
	var sum int64 = 0
	for {
		req, err := inStream.Recv()
		if err == io.EOF {
			resp := pb.SumManyResponse{
				Op1: sum,
			}
			return inStream.SendAndClose(&resp)
		}
		if err != nil {
			log.Fatalf("failed to read stream request for SumMany %v", err)
		}
		log.Printf("Received request %v.\n", req.Op1)
		sum += req.Op1
	}
}

// Cumulative sum receives a stream of int64 to sum and responds to each request with the sum so far
// The signature of this method must match the equivalent in the generated CalculatorServiceServer interface so our server is implementing the handler method
func (s *CalcServer) CumulativeSum(inStream grpc.BidiStreamingServer[pb.CumulativeSumRequest, pb.CumulativeSumResponse]) error {

	log.Println("Cumulative sum called")
	var sum int64 = 0
	for {
		req, err := inStream.Recv()
		if err == io.EOF {
			return nil // we don't call inStream.SendAndClose because the server is sending multiple responses
		}
		if err != nil {
			log.Fatalf("failed to read stream request for CumulativeSum %v", err)
		}
		log.Printf("\nReceived CumulativeSum request to add %d to %d", req.Input, sum)
		sum += req.Input
		resp := pb.CumulativeSumResponse{
			Result: sum,
		}
		// Note: the server doesn't know how many request the client is sending
		log.Printf("Sending sum so far %d", sum)
		err = inStream.Send(&resp)
		if err != nil {
			log.Fatalf("failed to send CumulativeSumResponse to stream.Error: %v", err)
		}
	}
}

// CountDown receives a single int64 and returns all the values from counting down
// from supplied value to zero as a stream
// The signature of this method must match the equivalent in the generated CalculatorServiceServer interface so our server is implementing the handler method
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

// countDown returns a slice containing all the values from n-1
// counting down to zero
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

// SquareRoot returns the square root in string format of the supplied int64
// If a negative number is received then and error is returned
// The signature of this method must match the equivalent in the generated CalculatorServiceServer interface so our server is implementing the handler method
func (s *CalcServer) SquareRoot(ctx context.Context, in *pb.SqrRootRequest) (*pb.SqrRootResponse, error) {
	log.Printf("Received request %v.\n", in)
	if in.Input < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Input must be positive number. Received %d", in.Input))
	}
	resultF := math.Sqrt(float64(in.Input))
	result := strconv.FormatFloat(resultF, 'f', 10, 64)

	resp := pb.SqrRootResponse{
		Result: result,
	}
	log.Printf("Responding with %v.\n", resp.Result)
	return &resp, nil
}
