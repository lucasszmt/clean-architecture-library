package register_user

import (
	user2 "awesomeLibraryProject/domain/userctx"
)

type RegisterUser struct {
	UserRepository user2.Repository
}

func (r *RegisterUser) Execute(data RegisterUserDto) {
	panic("implement me")
}
