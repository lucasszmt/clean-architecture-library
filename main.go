package main

import (
	"awesomeLibraryProject/database"
	"awesomeLibraryProject/infra/repository"
	"fmt"
	"log"
)

func main() {
	err := database.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	repo := repository.NewUserPostgres(database.Db)
	users, err := repo.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)
}
