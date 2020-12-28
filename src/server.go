package main

import (
	"api-message-security/src/hmac"
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
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		c.JSON(200, gin.H{
			"message": "pong",
			"body":    string(body),
		})
	})

	s.POST("/api/test", func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		c.JSON(200, gin.H{
			"message": "pong",
			"body":    string(body),
		})
	})
}

func (s *Server) Start() {
	if err := s.Run(); err != nil {
		panic(err)
	}
}
