package validator

import (
	"fmt"
	"testing"
)

func TestDo(t *testing.T) {
	e := employee{}

	err := validateStruct(e)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func TestDo2(t *testing.T) {
	e := employee{}
	err := validateStruct(e)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	e = employee{Name: "farhad"}
	err = validateStruct(e)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	e = employee{Name: "farhad", Age: 5}
	err = validateStruct(e)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	e = employee{Name: "farhad", Age: 25}
	err = validateStruct(e)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
