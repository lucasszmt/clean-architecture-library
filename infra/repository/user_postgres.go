package repository

import (
	"awesomeLibraryProject/domain/library/user"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type UserPostgres struct {
	db *sql.DB
}

type UserModel struct {
	id        int
	name      string
	email     string
	createdAt time.Time
}

func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (u *UserPostgres) GetAll() (users []*user.User, err error) {
	rows, err := u.db.Query("SELECT users.id, users.name, users.email, users.created_at FROM users")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		u := UserModel{}
		err = rows.Scan(&u.id, &u.name, &u.email, &u.createdAt)
		if err != nil {
			return
		}
		log.Println(u.createdAt)
		entity, _ := user.NewPresentationUser(u.id, u.name, u.email, u.createdAt)
		users = append(users, entity)
	}
	return
}

func (u *UserPostgres) FindById(id int) (*user.User, error) {
	panic("implement me")
}

func (u *UserPostgres) Delete(user *user.User) error {
	panic("implement me")
}

func (u *UserPostgres) Update(user *user.User) error {
	panic("implement me")
}

func (u *UserPostgres) Register(user *user.User) error {
	smt, err := u.db.Prepare(
		"INSERT INTO users(name, email, password, created_at) VALUES($1, $2, $3, $4);")
	if err != nil {
		log.Fatal("Sql Err:", err)
	}
	_, errExec := smt.Exec(user.GetName(), user.GetEmail(), user.GetPassword(), user.CreatedAt())
	if errExec != nil {
		log.Fatal(errExec)
	}

	return nil
}
