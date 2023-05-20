package main

import "fmt"

func main() {

	fmt.Println("enter somthing: ")
	var first string
	fmt.Scanln(&first)
	fmt.Print("Input: " + first + "\n")
}
