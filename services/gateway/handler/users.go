package handler

import (
	"context"
	"fmt"
	"log"
	"micr-go/core/heplers"
	"micr-go/services/gateway/validator"
	usersv "micr-go/services/users/pb"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

func FineOneUser(c *gin.Context) {
	conn := user.connGrpc()
	defer conn.Close()
	u := usersv.NewUsersClient(conn)
	email, id := c.Query("email"), c.Query("id")

	resp, err := u.FineOne(context.Background(), &usersv.GetOneRequest{
		Id:    id,
		Email: email,
	})

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"status": false, "error": err})
		return
	}

	c.JSON(200, resp)
}

func CreateUser(c *gin.Context) {
	var v validator.CreateUser
	if errValid := c.ShouldBindWith(&v, binding.Form); errValid != nil {
		c.AbortWithStatusJSON(400, gin.H{"status": false, "error": errValid})
		return
	}

	conn := user.connGrpc()
	defer conn.Close()
	u := usersv.NewUsersClient(conn)

	req := usersv.User{
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
		Email:    c.PostForm("email"),
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

func LoginUser(c *gin.Context) {
	var v validator.CreateUser
	if errValid := c.ShouldBindWith(&v, binding.Form); errValid != nil {
		c.AbortWithStatusJSON(400, gin.H{"status": false, "error": errValid})
		return
	}

	conn := user.connGrpc()
	defer conn.Close()
	u := usersv.NewUsersClient(conn)
	res, err := u.LoginUser(context.Background(), &usersv.LoginRequest{
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	})

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"status": false, "error": err})
		return
	}

	c.JSON(200, res)
}
