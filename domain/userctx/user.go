package userctx

import (
	"awesomeLibraryProject/domain/bookctx"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	id        string
	name      string
	password  string
	email     *Email
	books     []bookctx.Book
	createdAt time.Time
	updatedAt time.Time
}

func NewUser(name string, password string, email string) (*User, error) {
	e, errEmail := NewEmail(email)
	if errEmail != nil {
		return nil, errEmail
	}
	user := User{
		name:      name,
		password:  password,
		email:     e,
		createdAt: time.Now(),
	}

	hashedPassword, passErr := generatePassword(password)
	if passErr != nil {
		return nil, passErr
	}
	user.password = hashedPassword
	errValidate := user.Validate()
	if errValidate != nil {
		return nil, ErrInvalidUser
	}
	return &user, nil
}

func NewPresentationUser(id string, name string, email string) (user *User, err error) {
	e, err := NewEmail(email)
	if err != nil {
		return nil, err
	}
	user = &User{id: id, name: name, email: e}
	return
}

func (u *User) GetId() string {
	return u.id
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) Validate() error {
	if u.name == "" || u.password == "" {
		return ErrInvalidUser
	}
	return nil
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEmail() string {
	return u.email.String()
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func generatePassword(rawPassword string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 10)
	if err != nil {
		return "", err
	}
	return string(password), nil
}

func (u *User) String() string {
	return fmt.Sprintf("ID:%d Name: %s, Email: %s, Created At: %s", u.id, u.name, u.email, u.createdAt)
}
