package main

import (
	"fmt"
)

func main() {
	students := map[string]string{
		"cap": "max",
		"ass": "ted",
	}
	fmt.Println(students["cap"])
	students2 := make(map[string]string)
	students2["cap2"] = "max2"
	students2["ass2"] = "ted2"
	fmt.Println(students2["cap2"])

}
