package main

import (
	"context"
	"log"

	pb "github.com/judewood/gRPCSample/blog/proto"
)

func CreateBlog(c pb.BlogServiceClient) {

	log.Println("requesting create blog")
	// call the generated client function for this endpoint
	resp, err := c.CreateBlog(context.Background(), &pb.Blog{
		AuthorId: "1237777",
		Title: "Dial M for Murder",
		Content: "Always watch your back...",
	})
	if err != nil {
		log.Fatalf("failed to request Sum. Error: %v", err)
	}
	log.Printf("Result id: %s ",  resp.Id)
}
