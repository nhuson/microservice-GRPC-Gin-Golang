package main

import (
	"context"
	"fmt"
	"log"
	"micr-go/core/db"
	"micr-go/services/posts/handler"
	pb "micr-go/services/posts/pb"
	"net"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
)

func connectMongo() *mongo.Client {
	mongo := db.Connect("POST_MONGO_URL")
	err := mongo.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Println("Couldn't connect to the Mongo", err)
	} else {
		log.Println("Mongo Connected!")
	}

	return mongo
}

func main() {
	mongo := connectMongo()
	defer mongo.Disconnect(context.Background())
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatalf("Couldn't create connection tcp %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterPostsServer(srv, &handler.PostHandler{})
	log.Printf("Server start at port %s", os.Getenv("PORT"))

	srv.Serve(lis)
}
