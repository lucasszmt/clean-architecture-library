package main

import (
	"awesomeLibraryProject/application/user/register_user"
	"awesomeLibraryProject/infra/user"
	"fmt"
)

func main() {
	repo := &user.InMemoryUserRepository{}
	service := register_user.RegisterUser{UserRepository: repo}
	service.Execute(register_user.RegisterUserDto{
		Name: "Lucas Szeremeta",
		Cpf:  "08412535952",
	})
	fmt.Println(repo.Users)
}
