package heplers

import (
	"log"

	"github.com/gin-gonic/gin"
)

func GetDataHeader(c *gin.Context, key string) map[string]interface{} {
	data, error := c.Get(key)
	if !error {
		log.Printf("Can not get %s", key)
	}

	return data.(map[string]interface{})
}
