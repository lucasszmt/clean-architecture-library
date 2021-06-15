package main

import (
	"awesomeLibraryProject/config"
	"awesomeLibraryProject/database"
	"awesomeLibraryProject/domain/userctx"
	"log"
)

func main() {
	//Startup()
	//s := server.NewServer(":8000")
	//s.Run()
	u, _ := userctx.NewUser("Lucas Szeremeta", "123456789", "lucasszmt@gmail.com")
	log.Println(u.CheckPassword("123"))
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
