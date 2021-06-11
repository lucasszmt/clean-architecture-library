package controller

import (
	"awesomeLibraryProject/domain/userctx"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	Repo userctx.Repository
}

func (a *Auth) Login(c *gin.Context) {
	panic("Implement me")
}
