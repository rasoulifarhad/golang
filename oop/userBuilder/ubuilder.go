package userbuilder

import "fmt"

type User struct {
	Name      string
	Role      string
	MinSalary int
	MaxSalary int
}

type UserBuilder struct {
	User
}

func (ub *UserBuilder) Name(name string) *UserBuilder {
	ub.User.Name = name
	return ub
}

func (ub *UserBuilder) Role(role string) *UserBuilder {
	if role == "manager" {
		ub.User.MinSalary = 2000
		ub.User.MaxSalary = 4000
	}
	ub.User.Role = role

	return ub
}

func (ub *UserBuilder) Build() User {
	return ub.User
}

func main() {
	ub := &UserBuilder{}
	user := ub.
		Name("farhad rasouli").
		Role("manager").
		Build()

	fmt.Println(user)
}
