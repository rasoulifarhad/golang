package main

import (
	"fmt"
)

type people []string

func main() {
	var students = people{"max", "ted"}
	fmt.Println(students)

}
