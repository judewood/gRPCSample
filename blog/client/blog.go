package main

import (
	"context"
	"log"

	pb "github.com/judewood/gRPCSample/blog/proto"
)

func CreateBlog(c pb.BlogServiceClient) string {

	log.Println("requesting create blog")
	// call the generated client function for this endpoint
	resp, err := c.CreateBlog(context.Background(), &pb.Blog{
		AuthorId: "1237777",
		Title:    "Dial M for Murder",
		Content:  "Always watch your back...",
	})
	if err != nil {
		log.Fatalf("failed to request Sum. Error: %v", err)
	}
	log.Printf("Result id: %s ", resp.Id)
	return resp.Id
}

func GetBlog(c pb.BlogServiceClient, id string) {

	log.Printf("requesting  from id %s", id)
	// call the generated client function for this endpoint
	resp, err := c.GetBlog(context.Background(), &pb.BlogId{
		Id: id,
	})
	if err != nil {
		log.Fatalf("failed to request Sum. Error: %v", err)
	}
	log.Printf("Result  %v ", resp)
}
