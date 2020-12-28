package main

import (
	"api-message-security/src/hmac"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type Server struct {
	*gin.Engine
}

func NewServer() Server {
	r := gin.Default()
	r.Use(hmac.ValidGinMiddleware)
	return Server{r}
}

func (s *Server) RegisterRoute() {
	s.GET("/api/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	s.POST("/api/test", func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		var bodyJson interface{}
		json.Unmarshal(body, &bodyJson)

		c.JSON(200, gin.H{
			"message": "pong",
			"body":    bodyJson,
		})
	})
}

func (s *Server) Start() {
	if err := s.Run(); err != nil {
		panic(err)
	}
}
