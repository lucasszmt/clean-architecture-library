package main

import (
	"awesomeLibraryProject/config"
	"awesomeLibraryProject/database"
	"log"
)

func main() {
	Startup()
	//err := database.Init()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer database.Close()
	//err := config.Init()
	//if err != nil {
	//	log.Println(err)
	//}
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
