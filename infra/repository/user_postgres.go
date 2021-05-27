package repository

import (
	"awesomeLibraryProject/domain/library/user"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{db: db}
}

func (u User) FindById(id int) *user.User {
	panic("implement me")
}

func (u User) Delete(user *user.User) {
	panic("implement me")
}

func (u User) Update(user *user.User) {
	panic("implement me")
}

func (u User) Register(user *user.User) {
	smt, err := u.db.Prepare(
		"INSERT INTO users(name, email, password, created_at) VALUES($1, $2, $3, $4);")
	if err != nil {
		log.Fatal("Sql Err:", err)
	}
	result, errExec := smt.Exec(user.GetName(), user.GetEmail(), user.GetPassword(), user.CreatedAt())
	if errExec != nil {
		log.Fatal(errExec)
	}

	log.Println(result)
}
