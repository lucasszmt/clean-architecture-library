package server

import (
	"awesomeLibraryProject/app/http/controller"
	"awesomeLibraryProject/database"
	repo "awesomeLibraryProject/infra/repository"
)

func (s *Server) InitRoutes() {
	v1 := s.router.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			c := controller.Auth{Repo: repo.NewUserPostgres(database.Db)}
			auth.GET("/login", c.Login)
		}
	}
}
