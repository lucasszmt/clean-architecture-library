package user

import (
	"awesomeLibraryProject/domain/library/book"
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

func (u User) GetId() int {
	return u.id
}

func (u *User) SetEmail(email string) error {
	pattern, _ := regexp.Compile(`[A-Za-z]+@[a-z]+\.com(\.[a-z]+)?`)
	if pattern.MatchString(email) == true {
		return nil
	}
	return ErrInvalidEmail
}

func (u User) Validate() error {
	if u.name == "" || u.password == "" || u.email == "" {
		return ErrInvalidUser
	}
	return nil
}

func generatePassword(rawPassword string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 10)
	if err != nil {
		return "", err
	}
	return string(password), nil
}
