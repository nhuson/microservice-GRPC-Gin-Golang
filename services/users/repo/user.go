package repo

import (
	"context"
	"micr-go/core/db"
	pb "micr-go/services/users/pb"

	"go.mongodb.org/mongo-driver/mongo"
)

type User struct{}

func getInstance() *mongo.Collection {
	conn := db.Connect("USER_MONGO_URL")
	database := conn.Database("users")
	collection := database.Collection("user")

	return collection
}

func (u *User) CreateUser(user *pb.User) error {
	userModel := getInstance()
	_, err := userModel.InsertOne(context.Background(), user)

	return err
}
