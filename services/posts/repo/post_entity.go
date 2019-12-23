package repo

import "go.mongodb.org/mongo-driver/bson/primitive"

type PostItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Slug        string             `bson:"slug"`
	UserID      string             `bson:"user_id"`
	CreatedAt   string             `bson:"createdAt"`
	UpdatedAt   string             `bson:"updatedAt"`
}
