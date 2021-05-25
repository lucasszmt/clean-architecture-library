package user

type Repository interface {
	FindById(id int) *User
	Delete(user *User)
	Update(user *User)
	Register(user *User)
}
