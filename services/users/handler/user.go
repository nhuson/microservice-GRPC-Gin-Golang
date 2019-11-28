package handler

import (
	"context"
	"errors"
	"log"
	"micr-go/core/heplers"
	pb "micr-go/services/users/pb"
	"micr-go/services/users/repo"
	"time"

	"github.com/dgrijalva/jwt-go"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var user repo.User

type UsersHandler struct{}

func (u *UsersHandler) FindAll(ctx context.Context, req *pb.FindAllRequest) (*pb.FindAllResponse, error) {
	log.Printf("Receive message %s", req.Page)
	users := []*pb.User{
		{
			Username: "Son",
			Password: "123",
			Email:    "nhuson994@gmail.com",
			Address:  "Hanoi",
			Phone:    "09613435079",
		},
		{
			Username: "Steven",
			Password: "123",
			Email:    "steven94@gmail.com",
			Address:  "HaiDuong",
			Phone:    "09613435079",
		},
	}
	data := pb.FindAllResponse{
		Status:  true,
		Message: "Get Users",
		Data:    users,
	}
	return &data, nil
}

func (u *UsersHandler) CreateUser(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	log.Println("Create user.")
	userReq := req.GetUser()
	userFind := user.FindOne(ctx, bson.M{"email": userReq.GetEmail()})
	var uItem = repo.UserItem{}
	if error := userFind.Decode(&uItem); error == nil {
		return nil, errors.New("User already exist")
	}

	res, err := user.CreateUser(repo.UserItem{
		Email:    userReq.GetEmail(),
		Password: heplers.Encrypt(userReq.GetPassword()),
	})

	token := heplers.GenerateToken(&jwt.MapClaims{
		"userId": res.InsertedID.(primitive.ObjectID).Hex(),
		"email":  userReq.GetEmail(),
		"exp":    time.Now().Add(720 * time.Hour).Unix(),
	})

	data := pb.CreateResponse{
		Status:  true,
		Message: "Create user successfully!",
		Token:   token,
	}

	if err != nil {
		return nil, err
	}

	return &data, nil
}
