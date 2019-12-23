package repo

import (
	"context"
	"micr-go/core/db"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

type Post struct{}

func getInstance() *mongo.Collection {
	conn := db.Connect("POST_MONGO_URL").Database("posts").Collection("posts")
	return conn
}

func (p *Post) CreatePost(ctx context.Context, post PostItem) (*mongo.InsertOneResult, error) {
	postModel := getInstance()
	res, err := postModel.InsertOne(ctx, post)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *Post) ListPost(ctx context.Context) (*mongo.Cursor, error) {
	postModel := getInstance()
	cursor, err := postModel.Find(ctx, bson.M{})
	return cursor, err
}
