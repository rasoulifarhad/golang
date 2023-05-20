package main

import (
	"fmt"
)

func main() {
	var students = []string{"Max", "anna"}
	students = append(students, "ted")
	fmt.Println(students)

	s := make([]string, 3)
	fmt.Println("emp: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("s: ", s)
	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e")
	fmt.Println("apd: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	l := s[2:5]
	fmt.Println("slc1: ", l)

	l = s[:5]
	fmt.Println("slc2: ", l)

	l = s[2:]
	fmt.Println("slc3: ", l)

}
