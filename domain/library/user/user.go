package user

import "fmt"

type User struct {
	id       int
	name     string
	email    string
	password string
	phone    string
	cpf      string
	address  string
}

func NewUser(id int, name string, email string, password string, phone string, cpf string, address string) *User {
	return &User{id: id, name: name, email: email, password: password, phone: phone, cpf: cpf, address: address}
}

func NewUserWithNameCPF(name string, cpf string) *User {
	return &User{name: name, cpf: cpf}
}

func (u User) GetId() int {
	return u.id
}

func (u User) String() string {
	return fmt.Sprintf("Name: %s, Email: %s, CPF:%s", u.name, u.email, u.cpf)
}
