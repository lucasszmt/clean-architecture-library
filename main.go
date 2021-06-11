package main

import (
	"awesomeLibraryProject/config"
	"awesomeLibraryProject/database"
	"awesomeLibraryProject/domain/userctx"
	"awesomeLibraryProject/infra/repository"
	"log"
)

func main() {
	Startup()
	u, _ := userctx.NewUser("Lucas", "123456", "lucas@gmail.com")
	err := repository.UserPostgresRepo.Insert(u)
	log.Println(err)
}

func Startup() {
	err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	if err = database.Init(); err != nil {
		log.Fatal(err)
	}
}
