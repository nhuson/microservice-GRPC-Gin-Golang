package handler

import (
	"context"
	"log"
	pb "micr-go/services/users/pb"
	"micr-go/services/users/repo"
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

	err := user.CreateUser(req.User)
	data := pb.CreateResponse{
		Status:  true,
		Message: "Create user successfully!",
	}
	if err != nil {
		return nil, err
	}

	return &data, nil
}
