package routers

import (
	"fmt"
	"micr-go/services/gateway/handler"
	"micr-go/services/gateway/middleware"
	"os"

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
		userR.GET("/:id", handler.FineUserById)
		userR.PATCH("/:id", handler.UpdateUser)
	}

	postR := router.Group("/api/v1/posts")
	{
		postR.GET("", handler.ListPost)
		postR.POST("", handler.CreatePost)
	}

	router.Run(fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT")))

	return router
}
