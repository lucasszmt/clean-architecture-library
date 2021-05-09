package register_user

import "awesomeLibraryProject/domain/library/user"

type RegisterUser struct {
	UserRepository user.Repository
}

func (r *RegisterUser) Execute(data RegisterUserDto) {
	userWithNameCPF := user.NewUserWithNameCPF(data.Name, data.Cpf)
	r.UserRepository.Register(userWithNameCPF)
}
