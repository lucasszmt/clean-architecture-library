package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Init() error {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	type prs struct {
		name string
		cpf  string
	}
	p := &prs{}
	r := db.QueryRow("SELECT users.name, users.cpf FROM users").Scan(p)
	fmt.Println(r)
	fmt.Println(p)
	return nil
}
