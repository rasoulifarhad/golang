package main

import "fmt"

func main() {
	// An example which shows how to use type assertions(asserted types are non-interface types):
	// Compiler will deduce the type of 123 as int.
	var x interface{} = 123

	// Case 1:
	n, ok := x.(int)
	fmt.Println(n, ok) // 123 true
	n = x.(int)
	fmt.Println(n) // 123

	// Case 2:
	a, ok := x.(float64)
	fmt.Println(a, ok) // 0 false

	// Case 3:
	//a = x.(float64) // will panic

	// Another example which shows how to use type assertions (asserted types are interface types):
	var x1 interface{} = DummyWriter{}
	var y interface{} = "abc"
	// Now the dynamic type of y is "string".
	var w Writer
	var ok1 bool

	// Type DummyWriter implements both
	// Writer and interface{}.
	w, ok1 = x1.(Writer)
	fmt.Println(w, ok1) // {} true
	x1, ok1 = w.(interface{})
	fmt.Println(x1, ok1) // {} true

	// The dynamic type of y is "string",
	// which doesn't implement Writer.
	w, ok1 = y.(Writer)
	fmt.Println(w, ok1) //  false
	//w = y.(Writer)      // will panic

	//Here is an example in which a type-switch control flow code block is used.

	values := []interface{}{
		456, "abc", true, 0.33, int32(789),
		[]int{1, 2, 3}, map[int]bool{}, nil,
	}
	for _, x := range values {
		// Here, v is declared once, but it denotes
		// different variables in different branches.
		switch v := x.(type) {
		case []int: // a type literal
			// The type of v is "[]int" in this branch.
			fmt.Println("int slice:", v)
		case string: // one type name
			// The type of v is "string" in this branch.
			fmt.Println("string:", v)
		case int, float64, int32: // multiple type names
			// The type of v is "interface{}",
			// the same as x in this branch.
			fmt.Println("number:", v)
		case nil:
			// The type of v is "interface{}",
			// the same as x in this branch.
			fmt.Println(v)
		default:
			// The type of v is "interface{}",
			// the same as x in this branch.
			fmt.Println("others:", v)
		}
		// Note, each variable denoted by v in the
		// last three branches is a copy of x.
	}

}

type Writer interface {
	Write(buf []byte) (int, error)
}

type DummyWriter struct{}

func (DummyWriter) Write(buf []byte) (int, error) {
	return len(buf), nil
}
