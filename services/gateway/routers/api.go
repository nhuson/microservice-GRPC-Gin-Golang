package routers

import (
	"fmt"
	"micr-go/core/heplers"
	"micr-go/services/gateway/handler"
	"micr-go/services/gateway/middleware"

	"github.com/gin-gonic/gin"
)

func Api() *gin.Engine {
	router := gin.Default()
	/* Public API */
	router.GET("/api/v1/ping", handler.Pong)
	router.POST("/api/v1/signup", handler.CreateUser)
	router.POST("/api/v1/login", handler.LoginUser)

	/* Private API */
	router.Use(middleware.Authorization())
	userR := router.Group("/api/v1/users")
	{
		userR.GET("", handler.FindAllUser)
	}

	router.Run(fmt.Sprintf("0.0.0.0:%s", heplers.Getenv("PORT")))

	return router
}
