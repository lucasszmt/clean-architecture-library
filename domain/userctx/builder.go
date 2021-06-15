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

func (u *UserBuilder) Email(email string) *UserBuilder {
	e, _ := NewEmail(email)
	u.user.email = e
	return u
}

func (u *UserBuilder) Password(password string) *UserBuilder {
	u.user.password = password
	return u
}

func (u *UserBuilder) Build() *User {
	return &u.user
}
