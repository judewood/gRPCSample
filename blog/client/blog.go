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
		log.Fatalf("failed to create blog. Error: %v", err)
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
		log.Fatalf("failed to get blog. Error: %v", err)
	}
	log.Printf("Result  %v ", resp)
}

func UpdateBlog(c pb.BlogServiceClient, id string) error {

	log.Println("requesting update blog")
	// call the generated client function for this endpoint
	_, err := c.UpdateBlog(context.Background(), &pb.Blog{
		Id: id,
		AuthorId: "666",
		Title:    "Swallows and Amazons",
		Content:  "One fine day...",
	})
	if err != nil {
		log.Printf("err response %v", err)
		return err
	}
	log.Printf("updated id %s OK", id)
	return nil
}

func DeleteBlog(c pb.BlogServiceClient, id string) {
	log.Printf("requesting delete of blog Id: %s", id)
	// call the generated client function for this endpoint
	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{
		Id: id,
	})
	if err != nil {
		log.Fatalf("failed to delete blog. Error: %v", err)
	}
	log.Printf("Deleted blog id : %s", id)
}
