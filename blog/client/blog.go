package main

import (
	"context"
	"io"
	"log"

	pb "github.com/judewood/gRPCSample/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func CreateBlog(c pb.BlogServiceClient, blog *pb.Blog) string {
	resp, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("failed to create blog. Error: %v\n", err)
	}
	log.Printf("Blog id: %s created\n", resp.Id)
	return resp.Id
}

func GetBlog(c pb.BlogServiceClient, id string) {
	resp, err := c.GetBlog(context.Background(), &pb.BlogId{
		Id: id,
	})
	if err != nil {
		log.Fatalf("failed to get blog. Error: %v\n", err)
	}
	log.Printf("Got blog: %v\n", resp)
}

func UpdateBlog(c pb.BlogServiceClient, blog *pb.Blog) error {
	_, err := c.UpdateBlog(context.Background(), blog)
	if err != nil {
		log.Printf("failed to update blog.Error: %v\n", err)
		return err
	}
	log.Printf("Updated blog: %v\n", blog)
	return nil
}

func DeleteBlog(c pb.BlogServiceClient, id string) {
	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{
		Id: id,
	})
	if err != nil {
		log.Fatalf("failed to delete blog id %s. Error: %v", id, err)
	}
	log.Printf("Deleted blog id : %s", id)
}

func ListBlog(c pb.BlogServiceClient) {
	stream, err := c.ListBlog(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("failed to get stream for ListBlog requests. Error %v\n", err)
	}
	counter := 0
	for {
		msg, err := stream.Recv()
		if err == io.EOF { //stream closed by server
			log.Println("end of blog list")
			break
		}
		if err != nil {
			log.Fatalf("failed to list all blogs. Error %v\n", err)
		}
		counter++
		log.Printf("Blog: %d :  %v", counter, msg)
	}
}
