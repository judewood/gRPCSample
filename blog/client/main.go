package main

import (
	"log"

	pb "github.com/judewood/gRPCSample/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var endpointUrl = "localhost:4444"

func main() {
	opts := []grpc.DialOption{}
	tls := true //true to use SSL  - must match server setting
	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "") //empty string because we are using localhost
		if err != nil {
			log.Fatalf("failed to get client ssl credentials. Error %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	// create a persistent connection to a server endpoint
	conn, err := grpc.NewClient(endpointUrl, opts...)
	if err != nil {
		log.Fatalf("failed to create connection. Error: %v", err)
	}
	//ensure the connection is closed before its enclosing function returns
	defer conn.Close()

	// create an concrete  client struct from the generated code
	c := pb.NewBlogServiceClient(conn)

	id := CreateBlog(c)
	GetBlog(c, id)
	UpdateBlog(c,id)
	GetBlog(c, id)
}
