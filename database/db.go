package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func Init() (err error) {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	Db, err = sql.Open("postgres", dsn)
	return err
}

func Close() {
	Db.Close()
}
