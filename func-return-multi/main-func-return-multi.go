package main

import (
	"fmt"
)

func main() {

	fmt.Println(person())
	name, age := person()
	fmt.Println(name)
	fmt.Println(age)

}

func person() (string, int) {

	return "Max", 20

}
