package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/judewood/gRPCSample/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Sum sends two int64s to the server as a single (unary) request and expects a unary  response from the server
// The response is logged out
func Sum(c pb.CalculatorServiceClient, op1, op2 int64) {
	log.Printf("requesting sum of %d and %d", op1, op2)
	// call the generated client function for this endpoint
	resp, err := c.Sum(context.Background(), &pb.SumRequest{
		Op1: op1, Op2: op2,
	})
	if err != nil {
		log.Fatalf("failed to request Sum. Error: %v", err)
	}
	log.Printf("Result of %d + %d is %d", op1, op2, resp.Result)
}

// Sum sends two int64s to the server as a single (unary) request and expects a unary  response from the server
// The response is logged out
func SumDelay(c pb.CalculatorServiceClient, op1, op2 int64, deadline time.Duration) {
	log.Printf("requesting sum of %d and %d with deadline %#v seconds", op1, op2, deadline/1000000000)
	// call the generated client function for this endpoint
	ctx, cancelFunc := context.WithTimeout(context.Background(), deadline)
	defer cancelFunc()

	resp, err := c.SumDelay(ctx, &pb.SumRequest{
		Op1: op1, Op2: op2,
	})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Message from server %s\n", e.Message())
			log.Printf("Status code from server %s\n", e.Code())
			if e.Code() == codes.InvalidArgument {
				log.Println("Try giving the server more time")
			}
		} else {
			log.Fatalf("failed to request Sum. Error: %v", err)
		}
	} else {
		log.Printf("Result of %d + %d is %d", op1, op2, resp.Result)
	}
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

func CumulativeSum(c pb.CalculatorServiceClient, in []int64) {
	stream, err := c.CumulativeSum(context.Background())
	if err != nil {
		log.Fatalf("failed to get stream for CumulativeSum requests")
	}

	reqs := make([]pb.CumulativeSumRequest, len(in))
	for i := range in {
		reqs[i] = pb.CumulativeSumRequest{
			Input: in[i],
		}
	}

	// create a signally channel so we can block exit until all responses are received
	waitCh := make(chan struct{})

	//send the requests in this go routine
	go func() {
		for i := range reqs {
			i := i
			log.Printf("\n Sending CumulativeSumRequest req to server. %d ", reqs[i].Input)
			err := stream.Send(&reqs[i])
			if err != nil {
				log.Fatalf("failed to get stream for CumulativeSum requests")
			}
		}
		stream.CloseSend()
	}()

	// get the responses in this go routine
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break // happens after stream.CloseSend() is called in other go routine
			}
			if err != nil {
				log.Printf("failed to get resp for request : %v", err)
				break
			}
			log.Printf("Response sum : %d", resp.Result)
		}
		close(waitCh) //signals that we are done getting responses
	}()

	v, ok := <-waitCh //will block until the signaling channel is closed. Closing channel pushed something in so v, ok could be received
	// this line is just to show that we receive an empty value and Ok is false when channel is closed
	log.Printf("Channel closed with value = %v and OK = %v", v, ok)
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

// SquareRoot requests the square root of an integer and displays the response
func SquareRoot(c pb.CalculatorServiceClient, input int64) {
	log.Printf("requesting square root of %d", input)
	// call the generated client function for this endpoint
	resp, err := c.SquareRoot(context.Background(), &pb.SqrRootRequest{
		Input: input,
	})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Message from server %s\n", e.Message())
			log.Printf("Status code from server %s\n", e.Code())
			if e.Code() == codes.InvalidArgument {
				log.Println("Try using a positive input")
			}
			return
		} else {
			log.Fatalf("failed to request Square root . Error: %v", err)
		}
	}
	log.Printf("Square root of of %d is %s", input, resp.Result)
}
