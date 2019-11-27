package routers

import (
	"fmt"
	"micr-go/core/heplers"
	"micr-go/services/gateway/handler"

	"github.com/gin-gonic/gin"
)

func Api() *gin.Engine {
	router := gin.Default()
	router.GET("/api/v1/ping", handler.Pong)
	router.POST("/api/v1/signup", handler.CreateUser)

	userR := router.Group("/api/v1/users")
	{
		userR.GET("", handler.FindAllUser)
	}

	router.Run(fmt.Sprintf("0.0.0.0:%s", heplers.Getenv("PORT")))

	return router
}
