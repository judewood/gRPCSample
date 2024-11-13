package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/judewood/gRPCSample/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *BlogServer) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("CreateBlog invoked with %v", in)

	data := BlogItem{
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Printf("failed to create blog. Error %v", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error : %v", err),
		)
	}
	if createdId, ok := res.InsertedID.(primitive.ObjectID) ; !ok {
			return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot cast Id to string : %v", err),
		)	
	} else {
	blogId := pb.BlogId{
		Id: createdId.Hex(),
	}
	return &blogId, nil
	}


}

//func (s *BlogServer) GetBlog(context.Context, *BlogId) (*Blog, error)
//func (s *BlogServer) UpdateBlog(context.Context, *Blog) (*emptypb.Empty, error)
//func (s *BlogServer) DeleteBlog(context.Context, *BlogId) (*emptypb.Empty, error)
//func (s *BlogServer) ListBlog(*emptypb.Empty, grpc.ServerStreamingServer[Blog]) error
