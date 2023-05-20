package main

import "fmt"

type Filter interface {
	About() string
	Process([]int) []int
}

// UniqueFilter is used to remove duplicate numbers.
type UniqueFilter struct{}

func (UniqueFilter) About() string {
	return "remove duplicate numbers"
}
func (UniqueFilter) Process(inputs []int) []int {
	outs := make([]int, 0, len(inputs))
	pusheds := make(map[int]bool)
	for _, n := range inputs {
		if !pusheds[n] {
			pusheds[n] = true
			outs = append(outs, n)
		}
	}
	return outs
}

// MultipleFilter is used to keep only
// the numbers which are multiples of
// the MultipleFilter as an int value.
type MultipleFilter int

func (mf MultipleFilter) About() string {
	return fmt.Sprintf("keep multiples of %v", mf)
}
func (mf MultipleFilter) Process(inputs []int) []int {
	var outs = make([]int, 0, len(inputs))
	for _, n := range inputs {
		if n%int(mf) == 0 {
			outs = append(outs, n)
		}
	}
	return outs
}

// With the help of polymorphism, only one
// "filterAndPrint" function is needed.
func filterAndPrint(fltr Filter, unfiltered []int) []int {
	// Calling the methods of "fltr" will call the
	// methods of the value boxed in "fltr" actually.
	filtered := fltr.Process(unfiltered)
	fmt.Println(fltr.About()+":\n\t", filtered)
	return filtered
}

func main() {
	numbers := []int{12, 7, 21, 12, 12, 26, 25, 21, 30}
	fmt.Println("before filtering:\n\t", numbers)

	// Three non-interface values are boxed into
	// three Filter interface slice element values.
	filters := []Filter{
		UniqueFilter{},
		MultipleFilter(2),
		MultipleFilter(3),
	}

	// Each slice element will be assigned to the
	// local variable "fltr" (of interface type
	// Filter) one by one. The value boxed in each
	// element will also be copied into "fltr".
	for _, fltr := range filters {
		fmt.Println(fmt.Sprintf("(%v,%T)", fltr, fltr))
		numbers = filterAndPrint(fltr, numbers)
	}
}
