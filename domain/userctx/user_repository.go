package userctx

type Repository interface {
	GetAll() ([]*User, error)
	FindById(id int) (*User, error)
	Delete(id int) error
	Update(user *User) error
	Insert(user *User) error
}
