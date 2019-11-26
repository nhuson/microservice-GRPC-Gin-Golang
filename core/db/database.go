package db

import (
	"context"
	"log"
	"micr-go/core/heplers"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(url string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(heplers.Getenv(url)))

	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return client
}
