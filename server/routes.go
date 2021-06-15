package server

import (
	"awesomeLibraryProject/app/http/controller"
)

func (s *Server) InitRoutes() {
	v1 := s.router.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", controller.AuthController.Login)
			auth.GET("/user/:id", controller.AuthController.FindUser)
		}
	}
}
