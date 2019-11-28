package repo

import (
	"context"
	"micr-go/core/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct{}

func getInstance() *mongo.Collection {
	conn := db.Connect("USER_MONGO_URL").Database("users").Collection("users")
	return conn
}

func (u *User) CreateUser(user UserItem) (*mongo.InsertOneResult, error) {
	userModel := getInstance()
	res, err := userModel.InsertOne(context.Background(), user)

	return res, err
}

func (u *User) FindOne(ctx context.Context, options bson.M) *mongo.SingleResult {
	userModel := getInstance()
	resp := userModel.FindOne(ctx, options)

	return resp
}
