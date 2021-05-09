package user

type Repository interface {
	FindById(id int)
	Delete(user *User)
	Update(user *User)
	Register(user *User)
}
