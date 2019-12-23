package handler

import (
	"context"
	"log"
	"micr-go/core/heplers"
	"micr-go/services/gateway/validator"
	usersv "micr-go/services/users/pb"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"google.golang.org/grpc"
)

type UserHandler struct{}

var user UserHandler

func (u *UserHandler) connGrpc() *grpc.ClientConn {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(os.Getenv("USER_SERVICE"), grpc.WithInsecure())
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

func FineUserById(c *gin.Context) {
	conn := user.connGrpc()
	defer conn.Close()
	u := usersv.NewUsersClient(conn)
	id := c.Param("id")

	resp, err := u.FineOne(context.Background(), &usersv.GetOneRequest{
		Id: id,
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

func UpdateUser(c *gin.Context) {
	userH := heplers.GetDataHeader(c, "user")
	if userH["id"] != c.Param("id") {
		c.AbortWithStatusJSON(400, gin.H{"status": false, "error": "Not role to update this user."})
		return
	}

	if c.PostForm("email") != "" {
		c.AbortWithStatusJSON(400, gin.H{"status": false, "error": "Cannot update email"})
		return
	}

	if c.PostForm("password") != "" {
		c.AbortWithStatusJSON(400, gin.H{"status": false, "error": "Plz call api change password"})
		return
	}

	conn := user.connGrpc()
	defer conn.Close()
	u := usersv.NewUsersClient(conn)
	res, err := u.UpdateUser(context.Background(), &usersv.UpdateUserRequest{
		User: &usersv.User{
			Id:       c.Param("id"),
			Phone:    c.PostForm("phone"),
			Address:  c.PostForm("address"),
			Username: c.PostForm("username"),
		},
	})

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"status": false, "error": err})
		return
	}
	c.JSON(200, res)
}
