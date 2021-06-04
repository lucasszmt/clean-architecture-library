package user

type UseCase interface {
	GetUser()
	CreateUser()
	ListUsers()
	UpdateUser()
	DestroyUser()
}
