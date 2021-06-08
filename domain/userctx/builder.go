package userctx

type UserBuilder struct {
	user User
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{user: User{}}
}

func (u *UserBuilder) Id(id string) *UserBuilder {
	u.user.id = id
	return u
}

func (u *UserBuilder) Name(name string) *UserBuilder {
	u.user.name = name
	return u
}

func (u *UserBuilder) Email(email *Email) *UserBuilder {
	u.user.email = email
	return u
}

func (u *UserBuilder) Build() *User {
	return &u.user
}
