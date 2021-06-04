package database

import (
	"awesomeLibraryProject/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func Init() (err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Envs.DbHost, config.Envs.DbUser, config.Envs.DbPassword, config.Envs.DbName, config.Envs.DbPort)
	Db, err = sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	err = Db.Ping()
	return err
}

func Close() {
	Db.Close()
}
