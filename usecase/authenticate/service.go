package authenticate

import (
	"awesomeLibraryProject/domain/userctx"
	"awesomeLibraryProject/infra/repository"
)

var AuthService UseCase = Service{repo: repository.UserPostgresRepo}

type Service struct {
	repo userctx.Repository
}

func (s Service) Login(email string, password string) (string, error) {
	panic("implement me")
}

func (s Service) GenerateToken() error {
	panic("implement me")
}
