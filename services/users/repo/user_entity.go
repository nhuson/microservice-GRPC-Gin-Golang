package repo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserItem struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"`
	Email     string             `bson:"email"`
	Address   string             `bson:"address"`
	Phone     string             `bson:"phone"`
	CreatedAt string             `bson:"createdAt"`
	UpdatedAt string             `bson:"updatedAt"`
}
