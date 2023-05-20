package main

import (
	"fmt"
	"golang/example/generics/maps"
	"golang/example/generics/orderedmaps"
	"golang/example/generics/sets"
	"golang/example/generics/slices"
	"golang/example/generics/sorts"
	"golang/example/generics/metrics"
	"strings"
)

func main() {
	s := []int{1, 2, 3}

	floats := slices.Map(s, func(i int) float64 { return float64(i) })
	fmt.Println(floats)
	// Now floats is []float64{1.0, 2.0, 3.0}.

	sum := slices.Reduce(s, 0, func(i, j int) int { return i + j })
	fmt.Println(sum)

	// Now sum is 6.

	evens := slices.Filter(s, func(i int) bool { return i%2 == 0 })
	fmt.Println(evens)

	// Now evens is []int{2}.

	k := maps.Keys(map[int]int{1: 2, 2: 4})
	fmt.Println(k)

	// Create a set of ints.
	// We pass int as a type argument.
	// Then we write () because Make does not take any non-type arguments.
	// We have to pass an explicit type argument to Make.
	// Function argument type inference doesn't work because the
	// type argument to Make is only used for a result parameter type.
	set := sets.Make[int]()

	// Add the value 1 to the set s.
	set.Add(1)

	// Check that s does not contain the value 2.
	if set.Contains(2) {
		panic("unexpected 2")
	}
	if set.Contains(1) {
		fmt.Println("set contain 1")
	}

	//  sort package
	//s1 := []int32{3, 5, 2}
	//sort.OrderedSlice(s1)
	// Now s1 is []int32{2, 3, 5}

	//s2 := []string{"a", "c", "b"})
	//sort.OrderedSlice(s2)
	// Now s2 is []string{"a", "b", "c"}

	type Person struct {
		Name string
	}

	var per []*Person

	per = append(per, &Person{Name: "farhad"})
	per = append(per, &Person{Name: "amir"})
	per = append(per, &Person{Name: "bob"})

	// ...
	sorts.SliceFn(per, func(p1, p2 *Person) bool { return p1.Name < p2.Name })

	sli := slices.Append([]int{1, 2, 3}, 4, 5, 6)
	// Now s is []int{1, 2, 3, 4, 5, 6}.
	fmt.Println(sli)
	slices.Copy(sli[3:], []int{7, 8, 9})
	// Now s is []int{1, 2, 3, 7, 8, 9}
	fmt.Println(sli)
	
	//metrics
	F("one",1)
	F("one",1)
	F("one",1)
	F("one",1)
	F("one",1)
	F("two",2)
}

//orderedmaps

// Set m to an0 ordered map from string to string,
// using strings.Compare as the comparison function.
var m = orderedmaps.New[string , string](strings.Compare)

// Add adds the pair a, b to m.

func Add(a , b string) {
	m.Insert(a,b)
}

//metrics
var met = metrics.Metric2[string, int]{}

func F(s string, i int) {
	met.Add(s, i) // this call is type checked at compile time
}
