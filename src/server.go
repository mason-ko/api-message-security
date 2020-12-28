package main

import (
	"api-message-security/src/hmac"
	"github.com/gin-gonic/gin"
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
}

func (s *Server) Start() {
	if err := s.Run(); err != nil {
		panic(err)
	}
}
