package main

import (
	"log"

	pb "github.com/judewood/gRPCSample/blog/proto"
	"github.com/judewood/gRPCSample/internal/consts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	opts := []grpc.DialOption{}
	if consts.UseSSL {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "") //empty string because we are using localhost
		if err != nil {
			log.Fatalf("failed to get client ssl credentials. Error %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	// create a persistent connection to the server endpoint
	conn, err := grpc.NewClient(consts.ClientUrl, opts...)
	if err != nil {
		log.Fatalf("failed to create connection. Error: %v", err)
	}
	//ensure the connection is closed before its enclosing function returns
	defer conn.Close()

	// create an concrete client struct from the generated code
	c := pb.NewBlogServiceClient(conn)
	blogCrudOperations(c)
}

// blogCrudOperations exercises all our crud operations to simulate user activity
// USe tests instead in real app
func blogCrudOperations(c pb.BlogServiceClient) {
	blog1 := pb.Blog{
		AuthorId: "123",
		Title:    "Dial M for Murder",
		Content:  "Always watch your back...",
	}
	id1 := CreateBlog(c, &blog1)
	blog2 := pb.Blog{
		AuthorId: "456",
		Title:    "Swallows and Amazons",
		Content:  "On a fine day...",
	}
	GetBlog(c, id1)
	id2 := CreateBlog(c, &blog2)
	updatedBlog1 := pb.Blog{
		Id:       id1,
		AuthorId: "123",
		Title:    "Dial M for Murder",
		Content:  "On a dark scary night...",
	}
	UpdateBlog(c, &updatedBlog1)
	ListBlog(c)
	DeleteBlog(c, id2)
	ListBlog(c)
}
