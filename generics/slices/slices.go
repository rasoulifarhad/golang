// // Package slices implements various slice algorithms.
package slices

// Map turns a []T1 to a []T2 using a mapping function.
// This function has two type parameters, T1 and T2.
// This works with slices of any type.

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// Reduce reduces a []T1 to a single value using a reduction function.
func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

// Filter filters values from a slice using a filter function.
// It returns a new slice with only the elements of s
// for which f returned true.
func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r

}

// Here are some example calls of these functions. Type inference is used to determine the
// type arguments based on the types of the non-type arguments.

//s := []int{1, 2, 3}

//floats := slices.Map(s, func(i int) float64 { return float64(i) })
// Now floats is []float64{1.0, 2.0, 3.0}.

//sum := slices.Reduce(s, 0, func(i, j int) int { return i + j })
// Now sum is 6.

//evens := slices.Filter(s, func(i int) bool { return i%2 == 0 })
// Now evens is []int{2}.


// Append appends the contents of t to the end of s and returns the result.
// If s has enough capacity, it is extended in place; otherwise a
// new array is allocated and returned.	
func Append[T any] (s []T , t ...T) []T {
	lens := len(s)
	tot := lens + len(t)
	if tot < 0 {
		panic("Append: cap out of range")
	}
	if tot > cap(s) {
		news := make([]T ,tot, tot + tot/2)
		copy(news,s)
		s = news
	}
	s = s[:tot]
	copy(s[lens:], t)
	return s
}

// Copy copies values from t to s, stopping when either slice is
// full, returning the number of values copied.
func Copy[T any](s, t []T) int {
	i := 0
	for ; i < len(s) && i < len(t); i++ {
		s[i] = t[i]
	}
	return i
}