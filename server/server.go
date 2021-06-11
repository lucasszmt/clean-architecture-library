package server

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	host   string
	router *gin.Engine
}

func NewServer(host string) *Server {
	return &Server{host: host}
}

func (s *Server) Run() {
	s.router = gin.Default()
	s.InitRoutes()
	log.Fatal(s.router.Run(s.host))
}
