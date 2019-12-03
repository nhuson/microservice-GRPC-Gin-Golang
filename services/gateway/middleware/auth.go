package middleware

import (
	"context"
	"fmt"
	"log"
	pb "micr-go/services/users/pb"
	"os"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"
)

func findUserByEmail(email string) (*pb.GetOneResponse, error) {
	conn, err := grpc.Dial(os.Getenv("USER_SERVICE"), grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
	}

	defer conn.Close()
	u := pb.NewUsersClient(conn)

	resp, err := u.FineOne(context.Background(), &pb.GetOneRequest{Email: email})
	return resp, err
}

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(500, gin.H{"status": false, "error": "Missing authorization token!"})
			c.Abort()
			return
		}
		token = token[7:]
		mapClaim := jwt.MapClaims{}
		decode, error := jwt.ParseWithClaims(token, &mapClaim, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if !decode.Valid {
			c.JSON(500, gin.H{"status": false, "error": "Unauthorized!"})
			c.Abort()
			return
		}
		if error != nil {
			c.JSON(500, gin.H{"status": false, "error": error.Error()})
			c.Abort()
			return
		}

		user, err := findUserByEmail(fmt.Sprintf("%v", mapClaim["email"]))
		if err != nil {
			c.JSON(500, gin.H{"status": false, "error": "Unauthorized!"})
			c.Abort()
			return
		}

		if !user.GetStatus() {
			c.JSON(500, gin.H{"status": false, "error": "Unauthorized!"})
			c.Abort()
			return
		}

		c.Set("user", map[string]interface{}{
			"id":       user.Data.GetId(),
			"email":    user.Data.GetEmail(),
			"address":  user.Data.GetAddress(),
			"phone":    user.Data.GetPhone(),
			"username": user.Data.GetUsername(),
		})
		c.Next()
	}
}
