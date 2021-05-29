package main

import (
	"awesomeLibraryProject/database"
	"log"
)

func main() {
	err := database.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

}
