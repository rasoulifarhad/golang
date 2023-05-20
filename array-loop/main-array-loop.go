package main

import (
	"fmt"
)

func main() {
	var students = []string{"Max", "anna"}
	for i, value := range students {
		fmt.Println(i, value)
	}

	for _, value := range students {
		fmt.Println(value)
	}

}
