package main

import (
	"fmt"
)

type people []string

func (p people) print() {
	for _, card := range p {
		fmt.Println("student: " + card)
	}
}
func main() {
	var students = people{"max", "ted"}
	students.print()

}
