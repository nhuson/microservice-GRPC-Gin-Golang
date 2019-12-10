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

func (u *UsersHandler) CreateUser(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	userReq := req.GetUser()
	userFind := user.FindOne(ctx, bson.M{"email": userReq.GetEmail()})
	var uItem = repo.UserItem{}
	if error := userFind.Decode(&uItem); error == nil {
		return nil, errors.New("User already exist")
	}

	res, err := user.CreateUser(ctx, repo.UserItem{
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

func (u *UsersHandler) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.CreateResponse, error) {
	log.Println("Login method.")
	userFind := user.FindOne(ctx, bson.M{"email": req.GetEmail()})
	var uItem = repo.UserItem{}
	if error := userFind.Decode(&uItem); error != nil {
		return nil, errors.New("User not found")
	}

	comparePass := heplers.ComparePasswords(uItem.Password, req.GetPassword())
	if !comparePass {
		return nil, errors.New("Wrong email or password")
	}

	token := heplers.GenerateToken(&jwt.MapClaims{
		"userId":   uItem.ID.Hex(),
		"email":    uItem.Email,
		"username": uItem.Username,
		"address":  uItem.Address,
		"phone":    uItem.Phone,
		"exp":      time.Now().Add(720 * time.Hour).Unix(),
	})

	data := pb.CreateResponse{
		Status:  true,
		Message: "Login success!",
		Token:   token,
		User: &pb.User{
			Id:       uItem.ID.Hex(),
			Email:    uItem.Email,
			Username: uItem.Username,
			Address:  uItem.Address,
			Phone:    uItem.Phone,
		},
	}

	return &data, nil
}

func (u *UsersHandler) FindAll(ctx context.Context, req *pb.FindAllRequest) (*pb.FindAllResponse, error) {
	cursor, _ := user.FindAll(ctx)
	defer cursor.Close(ctx)
	userItem := repo.UserItem{}
	users := []*pb.User{}

	for cursor.Next(ctx) {
		if err := cursor.Decode(&userItem); err != nil {
			return nil, errors.New("Error decode users")
		}
		users = append(users, &pb.User{
			Id:       userItem.ID.Hex(),
			Username: userItem.Username,
			Email:    userItem.Email,
			Address:  userItem.Address,
			Phone:    userItem.Phone,
		})
	}

	data := pb.FindAllResponse{
		Status:  true,
		Message: "Get Users",
		Data:    users,
	}
	return &data, nil
}

func (u *UsersHandler) FineOne(ctx context.Context, req *pb.GetOneRequest) (*pb.GetOneResponse, error) {
	objID, _ := primitive.ObjectIDFromHex(req.GetId())
	userFind := user.FindOne(ctx, bson.M{"$or": []bson.M{
		bson.M{"_id": objID},
		bson.M{"email": req.GetEmail()},
	}})
	var uItem = repo.UserItem{}
	if error := userFind.Decode(&uItem); error != nil {
		return nil, errors.New("User cannot found")
	}

	data := pb.GetOneResponse{
		Status:  true,
		Message: "Get user",
		Data: &pb.User{
			Id:       uItem.ID.Hex(),
			Username: uItem.Username,
			Email:    uItem.Email,
			Phone:    uItem.Phone,
			Address:  uItem.Address,
		},
	}

	return &data, nil
}

func (u *UsersHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	userReq := req.GetUser()
	objID, _ := primitive.ObjectIDFromHex(userReq.GetId())
	array := map[string]interface{}{
		"address":  userReq.GetAddress(),
		"phone":    userReq.GetPhone(),
		"username": userReq.GetUsername(),
	}
	bsonUpdate := bson.M{}
	for k, v := range array {
		if v != "" {
			bsonUpdate[k] = v
		}
	}

	_, err := user.UpdateUser(ctx, bson.M{"_id": objID}, bson.M{"$set": bsonUpdate})

	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserResponse{
		Status:  true,
		Message: "Update user successfuly!",
	}, nil
}
