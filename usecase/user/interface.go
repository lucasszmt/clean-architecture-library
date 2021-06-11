package user

type UseCase interface {
	Authenticate()
	GetUser()
	CreateUser()
	ListUsers()
	UpdateUser()
	DestroyUser()
}
