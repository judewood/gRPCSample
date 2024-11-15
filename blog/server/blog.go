package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/judewood/gRPCSample/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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
	if createdId, ok := res.InsertedID.(primitive.ObjectID); !ok {
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

func (s *BlogServer) GetBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Printf("GetBlog invoked for id %s", in.Id)
	mongoId, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		log.Printf("failed to convert input to mongo Id format %v", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error when converting input to mongo Id format: %v", err),
		)
	}
	filter := bson.D{{Key: "_id", Value: mongoId}}
	var result BlogItem
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Printf("failed to decode result from mongoDb: Error: %v", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error when decoding result from mongoDb: %v", err),
		)
	}
	blog := documentToBlog(&result)
	log.Printf("Returning : %v", blog)
	return blog, nil
}

func (s *BlogServer) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	log.Printf("UpdateBlog invoked with %v", in)

	mongoId, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error cannot convert Id to mongo format: %v", err),
		)
	}

	data := &BlogItem{
		ID:       mongoId,
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	filter := bson.D{{Key: "_id", Value: mongoId}}
	res, err := collection.UpdateOne(ctx, filter, bson.M{"$set": data})
	if err != nil {
		log.Printf("failed to update blog. Error %v", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error : %v", err),
		)
	}
	if res.MatchedCount == 0 {
		log.Println("failed to update blog. Item not found")
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Internal error. Item not found : %v", err),
		)
	}
	return &emptypb.Empty{}, nil

}

func (s *BlogServer) DeleteBlog(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {
	log.Printf("DeleteBlog invoked with %s", in.Id)

	mongoId, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		log.Printf("failed to convert id to mongo format. Error %v", err)
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Invalid Id supplied: %s", in.Id))
	}
	res, err := collection.DeleteOne(ctx, bson.M{"_id": mongoId})

	if err != nil {
		log.Printf("failed to delete item id: %s. Error %v", in.Id, err)
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to delete item.  Internal error: %v", err))
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "Not found")
	}
	return &emptypb.Empty{}, nil
}

func (s *BlogServer) ListBlog(in *emptypb.Empty, stream grpc.ServerStreamingServer[pb.Blog]) error {
		log.Println("ListBlog invoked")
		
	ctx := context.Background()
	cursor,err := collection.Find(ctx, primitive.D{{}})
	if err != nil {
		log.Printf("failed to get results from mongoDb: Error: %v", err)
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error when get from mongoDb: %v", err),
		)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		data := BlogItem{}
		err:= cursor.Decode(&data)
		if err != nil {
			return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error when iterating ver mongoDb cursor: %v", err),
		)
		}
		log.Printf("\nReturning next blog: %v\n", &data)
		stream.Send(documentToBlog(&data))
	}
	if err = cursor.Err() ; err != nil {
			return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error when iterating ver mongoDb cursor: %v", err),
		)
	}
	return nil
}

func (s *BlogServer) SendCurrentTime(in *pb.InitiateCurrentTime, stream grpc.ServerStreamingServer[pb.CurrentTime]) error{
	interval := time.Duration(in.Interval)
	for {
		timeStr := time.Now().Format("2006-01-02 15:04:05")
		log.Printf("timer sending %s\n", timeStr)
		stream.Send(&pb.CurrentTime{CurrentTime: timeStr})
		time.Sleep(interval * time.Second)
	}
}
