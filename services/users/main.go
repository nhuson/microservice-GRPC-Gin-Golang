package main

import (
	"context"
	"fmt"
	"log"
	"micr-go/core/db"
	"micr-go/core/heplers"
	"micr-go/services/users/handler"
	pb "micr-go/services/users/pb"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
)

func connectMongo() *mongo.Client {
	mongo := db.Connect("USER_MONGO_URL")
	err := mongo.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the Mongo", err)
	} else {
		log.Println("Mongo Connected!")
	}

	return mongo
}

func main() {
	mongo := connectMongo()
	defer mongo.Disconnect(context.Background())

	/**
	* Open TCP to other instances can access by grpc
	 */
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", heplers.Getenv("PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := handler.UsersHandler{}
	srv := grpc.NewServer()
	pb.RegisterUsersServer(srv, &server)
	srv.Serve(lis)
}
