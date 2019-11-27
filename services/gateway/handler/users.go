package handler

import (
	"context"
	"fmt"
	"log"
	"micr-go/core/heplers"
	usersv "micr-go/services/users/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type UserHandler struct{}

var user UserHandler

func (u *UserHandler) connGrpc() *grpc.ClientConn {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":%s", heplers.Getenv("USER_SERVICE_PORT")), grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
	}

	return conn
}

func FindAllUser(c *gin.Context) {
	conn := user.connGrpc()
	defer conn.Close()

	u := usersv.NewUsersClient(conn)
	resp, err := u.FindAll(context.Background(), &usersv.FindAllRequest{Page: "1", PerPage: "2"})
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"status": false, "error": err})
		return
	}

	c.JSON(200, resp)
}

func CreateUser(c *gin.Context) {
	conn := user.connGrpc()
	defer conn.Close()
	u := usersv.NewUsersClient(conn)

	req := usersv.User{
		Username: "test",
		Password: "123123",
		Email:    "nhuson1094@gmail.com",
		Address:  "12",
		Phone:    "09231453",
	}

	resp, err := u.CreateUser(context.Background(), &usersv.CreateRequest{
		User: &req,
	})

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"status": false, "error": err})
		return
	}

	c.JSON(200, resp)
}
