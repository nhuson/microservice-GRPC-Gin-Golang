package handler

import (
	"context"
	"fmt"
	"log"
	"micr-go/core/heplers"
	"micr-go/services/gateway/validator"
	postser "micr-go/services/posts/pb"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"google.golang.org/grpc"
)

type PostHandler struct{}

var post PostHandler

func (p *PostHandler) connGrpc() *grpc.ClientConn {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(os.Getenv("POST_SERVICE"), grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
	}

	return conn
}

func CreatePost(c *gin.Context) {
	var v validator.CreatePost
	if errValid := c.ShouldBindWith(&v, binding.Form); errValid != nil {
		c.AbortWithStatusJSON(400, gin.H{"status": false, "error": errValid})
		return
	}
	conn := post.connGrpc()
	defer conn.Close()
	p := postser.NewPostsClient(conn)

	user := heplers.GetDataHeader(c, "user")
	userID := fmt.Sprintf("%s", user["id"])
	res, err := p.CreatePost(context.Background(), &postser.CreatePostRequest{
		Post: &postser.Post{
			Title:       c.PostForm("title"),
			Description: c.PostForm("description"),
			UserId:      userID,
		},
	})
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"status": false, "error": err})
		return
	}
	c.JSON(200, res)
}

func ListPost(c *gin.Context) {
	conn := post.connGrpc()
	defer conn.Close()
	p := postser.NewPostsClient(conn)
	res, err := p.ListPost(context.Background(), &postser.ListPostRequest{})
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"status": false, "error": err})
		return
	}

	c.JSON(200, res)
}
