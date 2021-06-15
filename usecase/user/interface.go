package user

type UseCase interface {
	FindUserById(id string) (SimpleUserDTO, error)
	CreateUser()
	ListUsers()
	UpdateUser()
	DestroyUser()
}
