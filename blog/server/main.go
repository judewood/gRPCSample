package main

import (
	"context"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/judewood/gRPCSample/blog/proto"
)

var endpointUrl = "0.0.0.0:4444"
var collection *mongo.Collection

type BlogServer struct {
	pb.BlogServiceServer
}

func main() {

	connStr := "mongodb://root:password@localhost:27017"
	dbClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connStr))
	if err != nil {
		log.Fatalf("failed to get mongo client. Error: %v", err)
	}
	collection = dbClient.Database("blogdb").Collection("blog")

	log.Printf("connected to mongo with connection string %s", connStr)
	
	listener, err := net.Listen("tcp", endpointUrl)
	if err != nil {
		log.Fatalf("failed to get listener. Error: %v", err)
	}
	defer listener.Close()


	log.Printf("Listening on: %s", endpointUrl)

	opts := []grpc.ServerOption{}
	tls := true //true to use SSL  - must match client setting
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("failed to get ssl credentials. Error %v", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}
	s := grpc.NewServer(opts...)
	pb.RegisterBlogServiceServer(s, &BlogServer{})

	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("failed to start blog server . Error: %v", err)
	}
}
