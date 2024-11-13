package main

import (
	pb "github.com/judewood/gRPCSample/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

// documentToBlog maps Mongo item to protobuf Blog message
func documentToBlog(data *BlogItem) *pb.Blog {
	return &pb.Blog{
		Id:       data.ID.Hex(), //Hex converts ID to string
		AuthorId: data.AuthorID,
		Title:    data.Title,
		Content:  data.Content,
	}
}
