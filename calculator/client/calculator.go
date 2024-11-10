package main

import (
	"context"
	"io"
	"log"

	pb "github.com/judewood/gRPCSample/calculator/proto"
)

// Sum sends two int64s to the server as a single (unary) request and expects a unary  response from the server 
// The response is logged out
func Sum(c pb.CalculatorServiceClient, op1, op2 int64) {
	log.Printf("requesting sum of %d and %d", op1, op2)
	// call the generated client function for this endpoint
	resp, err := c.Sum(context.Background(), &pb.SumRequest{
		Op1: 1, Op2: 2,
	})
	if err != nil {
		log.Fatalf("failed to request Sum. Error: %v", err)
	}
	log.Printf("Result of %d + %d is %d", op1, op2, resp.Result)
}
// SumMany sends a slice int64s to the server as streamed requests and expects a unary  response from the server 
// The response is logged out
func SumMany(c pb.CalculatorServiceClient, in []int64) {
	log.Printf("summing: %#v", in)
	//create a slice of requests
	reqs := make([]pb.SumManyRequest, len(in))
	for i := range in {
		reqs[i] = pb.SumManyRequest{
			Op1: in[i],
		}
	}
	// get a stream from the generated code
	stream, err := c.SumMany(context.Background())
	if err != nil {
		log.Fatalf("failed to request SumMany with inputs %v", in)
	}
	// and use it to stream the requests to the server
	for i := range reqs {
		log.Printf("Sending %d for summing", reqs[i].Op1)
		stream.Send(&reqs[i])
	}
    // close the stream and get the response
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to get result from server. Error: %v", err)
	}
	// print the result calculated by the server
	log.Printf("Result of summing %v is %d", in, res.Op1)
}

// CountDown sends one int64 to the server as a single (unary) request and expects  stream of responses 
// in the form of a countdown to zero
// The response is logged out
func CountDown(c pb.CalculatorServiceClient, val int64) {
	log.Printf("requesting Countdown from %d", val)
	req := pb.CountDownRequest{Value: val}
	// send the request and get back a stream that will contain the responses
	stream, err := c.CountDown(context.Background(), &req)
	if err != nil {
		log.Fatalf("failed to send countdown  request. Error %v", err)
	}
	// loop until the server closes the stream
	for {
		msg, err := stream.Recv()
		if err == io.EOF { //stream closed by server
			break
		}
		if err != nil {
			log.Fatalf("failed during stream. Error %v", err)
		}
		//print out the latest countdown value received from the server
		log.Printf("received:  %d", msg.Count)
	}
}
