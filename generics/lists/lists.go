// Package lists provides a linked list of any type.
package lists

// List is a linked list.
type List[T any] struct {
	head , tail *element[T]
}

// An element is an entry in a linked list.
type element[T any] struct {
	val T
	next *element[T]
}

// Push pushes an element to the end of the list.
func (lst *List[T]) Push(v T){
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// Iterator ranges over a list.
type Iterator[T any] struct {
	next **element[T]
}

// Range returns an Iterator starting at the head of the list.
func (lst *List[T]) Range() *Iterator[T] {
	return &Iterator[T]{next: &lst.head}
}

// Next advances the iterator.
// It reports whether there are more elements.
func (it *Iterator[T]) Next() bool {
	if *it.next == nil {
		return false
	}
	it.next= &(*it.next).next
	return true
}

// Val returns the value of the current element.
// The bool result reports whether the value is valid.
func (it *Iterator[T]) Val() (T,bool) {
	if *it.next == nil {
		var zero T
		return zero , false
	}
	return (*it.next).val , true
}

// Transform runs a transform function on a list returning a new list.
func Transform[T1 , T2 any] (lst *List[T1] , f func(T1) T2) *List[T2] {
	ret  := &List[T2]{}
	it := lst.Range()
	for {
		if v , ok := it.Val(); ok {
			ret.Push(f(v))
		}
		if !it.Next(){
			break
		}

	}
	return ret
}

// NumericAbs matches numeric types with an Abs method.
type NumericAbs[T any] interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~complex64 | ~complex128
	Abs() T
}

// AbsDifference computes the absolute value of the difference of
// a and b, where the absolute value is determined by the Abs method.
func AbsDifference[T NumericAbs[T]](a, b T) T {
	d := a - b
	return d.Abs()
}

// OrderedNumeric matches numeric types that support the < operator.
type OrderedNumeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// Complex matches the two complex types, which do not have a < operator.
type Complex interface {
	~complex64 | ~complex128
}

type GenericSlice[T any] []T

// OrderedAbs is a helper type that defines an Abs method for
// ordered numeric types.
//type OrderedAbs[T OrderedNumeric] T

