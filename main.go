package main

import (
	"awesomeLibraryProject/database"
)

func main() {
	//url := "host=localhost user=postgres password=postgres dbname=postgres port=5432"
	//db, _ := gorm.Open(postgres.Open(url), &gorm.Config{})
	//err := db.AutoMigrate(&migrations.User{}, &migrations.Book{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	database.Init()

}
