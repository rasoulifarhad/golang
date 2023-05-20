package model

import (
	"fmt"
	"testing"
)

func TestDo(t *testing.T) {
	//STRUCTURE IDENTIFIER
	p := &Person{
		Name: "test",
		age:  21,
	}
	fmt.Println(p)
	c := &company{}
	fmt.Println(c)

	//STRUCTURE'S METHOD
	fmt.Println(p.GetAge())
	fmt.Println(p.getName())

	//STRUCTURE'S FIELDS
	fmt.Println(p.Name)
	fmt.Println(p.age)

	//FUNCTION
	person2 := GetPerson()
	fmt.Println(person2)
	companyName := getCompanyName()
	fmt.Println(companyName)

	//VARIBLES
	fmt.Println(companyLocation)
	fmt.Println(CompanyName)

}
