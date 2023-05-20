package govalidator

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

type employee struct {
	Name string `valid:"required"`
}

func validateStruct(e employee) error {
	_, err := govalidator.ValidateStruct(e)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	e := employee{}

	err := validateStruct(e)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
