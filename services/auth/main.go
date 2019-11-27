package main

import (
	"context"
	"fmt"
	"log"
	"micr-go/core/db"
	"micr-go/core/heplers"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", heplers.Getenv("PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
