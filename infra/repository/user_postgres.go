package repository

import (
	"awesomeLibraryProject/database"
	"awesomeLibraryProject/domain/userctx"
	_ "github.com/lib/pq"
	"log"
)

var UserPostgresRepo userctx.Repository = NewUserPostgres()

type UserPostgres struct{}

type UserModel struct {
	id       string
	name     string
	email    string
	password string
}

func NewUserPostgres() *UserPostgres {
	return &UserPostgres{}
}

func (u *UserPostgres) GetAll() (users []*userctx.User, err error) {
	rows, err := database.Db.Query("SELECT users.id, users.name, users.email FROM users")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		u := UserModel{}
		err = rows.Scan(&u.id, &u.name, &u.email)
		if err != nil {
			return
		}
		entity, _ := userctx.NewPresentationUser(u.id, u.name, u.email)
		users = append(users, entity)
	}
	return
}

func (u *UserPostgres) FindById(id int) (*userctx.User, error) {
	stmt, _ := database.Db.Prepare("SELECT id, name, email, password from users where users.id = $1")
	defer stmt.Close()
	usrModel := UserModel{}
	err := stmt.QueryRow(id).Scan(&usrModel.id, &usrModel.name, &usrModel.email, &usrModel.password)
	if err != nil {
		return nil, err
	}

	return userctx.NewPresentationUser(usrModel.id, usrModel.name, usrModel.email)
}

func (u *UserPostgres) Insert(user *userctx.User) error {
	stmt, err := database.Db.Prepare(
		"INSERT INTO users(name, email, password, created_at) VALUES($1, $2, $3, $4);")
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()
	res, errExec := stmt.Exec(user.GetName(), user.GetEmail(), user.GetPassword(), user.CreatedAt())
	if errExec != nil {
		log.Println(err)
		return errExec
	}
	lastId, err := res.RowsAffected()
	if err != nil {
		return err
	}
	rowCount, rowErr := res.RowsAffected()
	if rowErr != nil {
		return rowErr
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCount)
	return nil
}

func (u *UserPostgres) Delete(id int) error {
	stmt, err := database.Db.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}

func (u *UserPostgres) Update(user *userctx.User) error {
	stmt, err := database.Db.Prepare("UPDATE users SET name=$1, email=$2, password=$3 WHERE users.id = $4")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.GetName(), user.GetEmail(), user.GetPassword(), user.GetId())
	return err
}
