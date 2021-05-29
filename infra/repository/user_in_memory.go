package repository

import (
	user2 "awesomeLibraryProject/domain/userctx"
)

type InMemoryUserRepository struct {
	Users []user2.User
}

func (i *InMemoryUserRepository) FindById(id int) {
	panic("implement me")
}

func (i *InMemoryUserRepository) Delete(user *user2.User) {
	panic("implement me")
}

func (i *InMemoryUserRepository) Update(user *user2.User) {
	panic("implement me")
}

func (i *InMemoryUserRepository) Register(user *user2.User) {
	i.Users = append(i.Users, *user)
}

//
//func (i *InMemoryUserRepository) Update(repository *repository.UserPostgres) {
//	panic("implement me")
//}
//
//func (i *InMemoryUserRepository) FindById(id int) {
//	panic("implement me")
//}
//
//func (i *InMemoryUserRepository) Insert(repository repository.UserPostgres) {
//	i.Users = append(i.Users, repository)
//}
//
//func (i *InMemoryUserRepository) Delete(repository repository.UserPostgres) {
//	for index, elem := range i.Users {
//		if repository.GetId() == elem.GetId() {
//			i.Users = append(i.Users[:index], i.Users[index+1:]...)
//		}
//	}
//}
