package repo

import (
	"context"
	"errors"
	"micr-go/core/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct{}

func getInstance() *mongo.Collection {
	conn := db.Connect("USER_MONGO_URL").Database("users").Collection("users")
	return conn
}

func (u *User) CreateUser(ctx context.Context, user UserItem) (*mongo.InsertOneResult, error) {
	userModel := getInstance()
	res, err := userModel.InsertOne(ctx, user)

	return res, err
}

func (u *User) UpdateUser(ctx context.Context, filter bson.M, user bson.M) (*mongo.SingleResult, error) {
	userModel := getInstance()
	res := userModel.FindOneAndUpdate(ctx, filter, user, options.FindOneAndUpdate().SetReturnDocument(1))
	if res == nil {
		return nil, errors.New("Cannot update user")
	}
	return res, nil
}

func (u *User) FindOne(ctx context.Context, options bson.M) *mongo.SingleResult {
	userModel := getInstance()
	resp := userModel.FindOne(ctx, options)

	return resp
}

func (u *User) FindAll(ctx context.Context) (*mongo.Cursor, error) {
	userModel := getInstance()
	// count, err := userModel.CountDocuments(ctx, bson.M{})
	// var limit int64 = 2
	// var skip int64 = 1
	// findOpt := options.FindOptions{
	// 	Limit: &limit,
	// 	Skip:  &skip,
	// }
	cursor, err := userModel.Find(ctx, bson.M{})
	return cursor, err
}
