// Sample program to show how the program can't access an
// unexported identifier from another package.
package main

import (
	"fmt"
	"golang/example/unexport-identifiers/counters"
	"golang/example/unexport-identifiers/entities"
)

func main() {
	// Create a variable of the unexported type and initialize
	// the value to 10.
	//counter := counters.alertCounter(10)

	counter := counters.New(10)
	fmt.Printf("counter: %d\n", counter)

	//	u := entities.User{
	//		Name: "fsrhad",
	//		email: "ggggg",

	//	}

	//	fmt.Printf("User: %v\n",u)
	a := entities.Admin{
		Rights: 10,
	}
	a.Name = "ggggg"
	a.Email = "gffggggggggg"
	fmt.Printf("User: %v\n", a)
}
