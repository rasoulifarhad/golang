package validator

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

type employee struct {
	Name string `validate:"required"`
	Age  int    `validate:"required,gte=10,lte=20"`
}

func validateStruct(e employee) error {
	validate = validator.New()
	err := validate.Struct(e)

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
