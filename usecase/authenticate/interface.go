package authenticate

type UseCase interface {
	Login(email string, password string) (string, error)
	GenerateToken() error
}
