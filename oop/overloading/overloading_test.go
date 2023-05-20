package overloading

import (
	"fmt"
	"testing"
)

func TestDo(t *testing.T) {
	fmt.Println(add(1, 2))
	fmt.Println(add(1, 2, 3))
	fmt.Println(add(1, 2, 3, 4))
}

func TestParams(t *testing.T) {
	handle(1, "abc")
	handle("abc", "xyz", 3)
	handle(1, 2, 3, 4)
}

func TestPerson(t *testing.T) {
	err := addPerson("Tina", "Female", 20)
	if err != nil {
		fmt.Println("PersonAdd Error: " + err.Error())
	}

	err = addPerson("John", "Male")
	if err != nil {
		fmt.Println("PersonAdd Error: " + err.Error())
	}

	err = addPerson("Wick", 2, 3)
	if err != nil {
		fmt.Println("PersonAdd Error: " + err.Error())
	}
}
