package main

import (
	"api-message-security/src/hmac"
	"github.com/gin-gonic/gin"
)

// base64(hmac( method | url | timestamp ))

/*
curl -X GET "http://localhost:8080/api/test" \
-H "x_authorization: 91kKQixU8OZU1XcIBSSwo+G7l43M603De2LtjF/Khoo=" \
-H "x_timestamp: 2020-12-23T08:28:26.737Z"
 */

func main() {
	r := gin.Default()
	r.Use(hmac.ValidGinMiddleware)

	r.GET("/api/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}