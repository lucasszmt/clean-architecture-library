package user

import (
	"awesomeLibraryProject/domain/library/book"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"time"
)

type User struct {
	id        int
	name      string
	password  string
	email     string
	books     []book.Book
	createdAt time.Time
	updatedAt time.Time
}

func NewUser(name string, password string, email string) (*User, error) {
	user := User{
		name:      name,
		password:  password,
		createdAt: time.Now(),
	}
	if err := user.SetEmail(email); err != nil {
		return nil, err
	}
	hashedPassword, passErr := generatePassword(password)
	if passErr != nil {
		return nil, passErr
	}
	user.password = hashedPassword
	return &user, nil
}

func NewPresentationUser(id int, name string, email string, createdAt time.Time) (user *User, err error) {
	user = &User{id: id, name: name, email: email, createdAt: createdAt}
	return
}

func (u *User) GetId() int {
	return u.id
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) SetEmail(email string) error {
	pattern, _ := regexp.Compile(`[A-Za-z]+@[a-z]+\.com(\.[a-z]+)?`)
	if pattern.MatchString(email) == true {
		u.email = email
		return nil
	}
	return ErrInvalidEmail
}

func (u *User) Validate() error {
	if u.name == "" || u.password == "" || u.email == "" {
		return ErrInvalidUser
	}
	return nil
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEmail() string {
	return u.email
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
