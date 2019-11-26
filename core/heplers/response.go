package heplers

import (
	"github.com/gin-gonic/gin"
)

func ResponseSuccess(data interface{}, message string) interface{} {
	res := gin.H{
		"status":  true,
		"data":    data,
		"message": message,
	}
	return res
}

func ResponseError(code int, data interface{}, message string) interface{} {
	res := gin.H{
		"status":  false,
		"data":    data,
		"message": message,
	}

	return res
}
