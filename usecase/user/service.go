package user

import (
	"awesomeLibraryProject/domain/userctx"
	"awesomeLibraryProject/infra/repository"
)

var (
	UserService UseCase = Service{repository.UserPostgresRepo}
)

type Service struct {
	repo userctx.Repository
}

func (s Service) FindUserById(id string) (SimpleUserDTO, error) {
	u, _ := s.repo.FindById(id)
	dto := SimpleUserDTO{
		Id:    u.GetId(),
		Name:  u.GetName(),
		Email: u.GetEmail(),
	}
	return dto, nil
}

func (s Service) CreateUser() {
	panic("implement me")
}

func (s Service) ListUsers() {
	panic("implement me")
}

func (s Service) UpdateUser() {
	panic("implement me")
}

func (s Service) DestroyUser() {
	panic("implement me")
}
