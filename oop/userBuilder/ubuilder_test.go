package userbuilder

import (
	"fmt"
	"testing"
)

func TestDo(t *testing.T) {
	ub := &UserBuilder{}
	user := ub.
		Name("farhad rasouli").
		Role("manager").
		Build()

	fmt.Println(user)
}
