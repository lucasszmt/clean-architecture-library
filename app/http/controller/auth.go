package controller

import (
	"awesomeLibraryProject/usecase/authenticate"
	"awesomeLibraryProject/usecase/user"
	"github.com/gin-gonic/gin"
)

var AuthController = Auth{userService: user.UserService, authService: authenticate.AuthService}

type Auth struct {
	userService user.UseCase
	authService authenticate.UseCase
}

func (a *Auth) Login(c *gin.Context) {
	credentials := authenticate.AuthCredentialsDto{}
	err := c.ShouldBindJSON(&credentials)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Credentials"})
		return
	}

}

func (a *Auth) FindUser(c *gin.Context) {
	id := c.Param("id")
	userData, err := a.userService.FindUserById(id)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "user not found"})
		return
	}
	c.JSON(200, gin.H{"data": userData})
}
