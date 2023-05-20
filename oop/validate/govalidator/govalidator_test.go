package govalidator

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
