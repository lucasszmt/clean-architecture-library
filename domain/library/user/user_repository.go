package user

type Repository interface {
	GetAll() ([]*User, error)
	FindById(id int) (*User, error)
	Delete(user *User) error
	Update(user *User) error
	Register(user *User) error
}
