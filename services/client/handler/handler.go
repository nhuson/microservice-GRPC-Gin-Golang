package handler

import "github.com/gin-gonic/gin"

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": "pong",
	})
}
