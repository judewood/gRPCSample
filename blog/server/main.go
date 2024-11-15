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
	"github.com/judewood/gRPCSample/internal/consts"
)

// equivalent to a sql table
var collection *mongo.Collection

type BlogServer struct {
	pb.BlogServiceServer
}

func main() {
	ctx := context.Background()
	dbClient := getMongoDbClient(ctx)
	defer dbClient.Disconnect(ctx)

	if consts.UseHttp1 {
		server := SetupRouter()
		log.Println("HTTP1 Server started")
		server.Run(":4444")
	}
	
	listener := getListener()
	defer listener.Close()

	s := grpc.NewServer(getServerOptions()...)
	pb.RegisterBlogServiceServer(s, &BlogServer{})

	err := s.Serve(listener)
	if err != nil {
		log.Fatalf("failed to start blog server . Error: %v", err)
	}
}

// getServerOptions gets the options to configure our http2 server with
// currently only enables/disables SSL
func getServerOptions() []grpc.ServerOption {
	opts := []grpc.ServerOption{}
	if consts.UseSSL {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("failed to get ssl credentials. Error %v", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}
	return opts
}

// getListener returns a listener that monitors calls to our url
func getListener() net.Listener {
	listener, err := net.Listen(consts.TCP, consts.ServerUrl)
	if err != nil {
		log.Fatalf("failed to get listener. Error: %v", err)
	}
	log.Printf("Listening on: %s", consts.ServerUrl)
	return listener
}

// getMongoDbClient returns a client to connect to mongoDB
func getMongoDbClient(ctx context.Context) *mongo.Client {
	// in real app build this from config
	const connStr string = "mongodb://root:password@localhost:27017"
	dbClient, err := mongo.Connect(ctx, options.Client().ApplyURI(connStr))
	if err != nil {
		log.Fatalf("failed to connect to mongoDb. Error: %v\n", err)
	}
	collection = dbClient.Database("blogdb").Collection("blog")
	log.Printf("connected to mongo with connection string %s", connStr)
	return dbClient
}
