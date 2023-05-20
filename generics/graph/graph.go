package graph

// NodeConstraint is the type constraint for graph nodes:
// they must have an Edges method that returns the Edge's
// that connect to this Node.
type NodeConstraint[Edge any] interface {
	Edges() []Edge
}

// EdgeConstraint is the type constraint for graph edges:
// they must have a Nodes method that returns the two Nodes
// that this edge connects.
type EdgeConstraint[Node any] interface {
	Nodes() (from, to Node)
}

// Graph is a graph composed of nodes and edges.
type Graph[Node NodeConstraint[Edge], Edge EdgeConstraint[Node]] struct { ... }

// New returns a new graph given a list of nodes.
func New[Node NodeConstraint[Edge], Edge EdgeConstraint[Node]] (nodes []Node) *Graph[Node, Edge] {
	...
}

// ShortestPath returns the shortest path between two nodes,
// as a list of edges.
func (g *Graph[Node, Edge]) ShortestPath(from, to Node) []Edge { ... }


// Vertex is a node in a graph.
type Vertex struct { ... }

// Edges returns the edges connected to v.
func (v *Vertex) Edges() []*FromTo { ... }

// FromTo is an edge in a graph.
type FromTo struct { ... }

// Nodes returns the nodes that ft connects.
func (ft *FromTo) Nodes() (*Vertex, *Vertex) { ... }


// var g = graph.New[*Vertex, *FromTo]([]*Vertex{ ... })

// We could instantiate graph.Graph with the types NodeInterface and EdgeInterface, 
// since they implement the type constraints. There isn't much reason to instantiate a type 
// this way, but it is permitted.
// type NodeInterface interface { Edges() []EdgeInterface }
// type EdgeInterface interface { Nodes() (NodeInterface, NodeInterface) }


// Map calls the function f on every element of the slice s,
// returning a new slice of the results.
func Map[F, T any](s []F, f func(F) T) []T {
	r := make([]T, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// The two type parameters F and T are both used for input parameters, so function 
// argument type inference is possible. In the call
strs := Map([]int{1, 2, 3}, strconv.Itoa)

// To see the untyped constant rule in effect, consider:

// NewPair returns a pair of values of the same type.
func NewPair[F any](f1, f2 F) *Pair[F] { ... }

// In the call NewPair(1, 2) both arguments are untyped constants, so both are ignored 
// in the first pass. There is nothing to unify. We still have two untyped constants after 
// the first pass. Both are set to their default type, int. The second run of the type 
//unification pass unifies F with int, so the final call is NewPair[int](1, 2).

// In the call NewPair(1, int64(2)) the first argument is an untyped constant, so we ignore 
// it in the first pass. We then unify int64 with F. At this point the type parameter 
// corresponding to the untyped constant is fully determined, so the final call is 
// NewPair[int64](1, int64(2)).

// In the call NewPair(1, 2.5) both arguments are untyped constants, so we move on the 
// second pass. This time we set the first constant to int and the second to float64. 
//We then try to unify F with both int and float64, so unification fails, and we report 
// a compilation error.




// For an example of where constraint type inference is useful, let's consider a function 
// that takes a defined type that is a slice of numbers, and returns an instance of that 
// same defined type in which each number is doubled.

// It's easy to write a function similar to this if we ignore the defined type requirement.

// Double returns a new slice that contains all the elements of s, doubled.
func Double[E constraints.Integer](s []E) []E {
	r := make([]E, len(s))
	for i, v := range s {
		r[i] = v + v
	}
	return r
}

// bHowever, with that definition, if we call the function with a defined slice type, the 
// result will not be that defined type.

// MySlice is a slice of ints.
type MySlice []int

// The type of V1 will be []int, not MySlice.
// Here we are using function argument type inference,
// but not constraint type inference.
var V1 = Double(MySlice{1})

// We can do what we want by introducing a new type parameter.

// DoubleDefined returns a new slice that contains the elements of s,
// doubled, and also has the same type as s.
func DoubleDefined[S ~[]E, E constraints.Integer](s S) S {
	// Note that here we pass S to make, where above we passed []E.
	r := make(S, len(s))
	for i, v := range s {
		r[i] = v + v
	}
	return r
}

// Now if we use explicit type arguments, we can get the right type.

// The type of V2 will be MySlice.
var V2 = DoubleDefined[MySlice, int](MySlice{1})

// Function argument type inference by itself is not enough to infer the type arguments here, 
// because the type parameter E is not used for any input parameter. But a combination of 
// function argument type inference and constraint type inference works.

// The type of V3 will be MySlice.
var V3 = DoubleDefined(MySlice{1})

// First we apply function argument type inference. We see that the type of the argument is 
// MySlice. Function argument type inference matches the type parameter S with MySlice.

// We then move on to constraint type inference. We know one type argument, S. We see that 
// the type argument S has a structural type constraint.

// We create a mapping of known type arguments:

// {S -> MySlice}

// We then unify each type parameter with a structural constraint with the single type in 
// that constraint's type set. In this case the structural constraint is ~[]E which has the 
// structural type []E, so we unify S with []E. Since we already have a mapping for S, we 
// then unify []E with MySlice. As MySlice is defined as []int, that associates E with int. 
// We now have:

// {S -> MySlice, E -> int}

// We then substitute E with int, which changes nothing, and we are done. The type arguments 
//for this call to DoubleDefined are [MySlice, int].




// Consider this example of a function that expects a type T that has a Set(string) method 
// that initializes a value based on a string.

// Setter is a type constraint that requires that the type
// implement a Set method that sets the value from a string.
type Setter interface {
	Set(string)
}

// FromStrings takes a slice of strings and returns a slice of T,
// calling the Set method to set each returned value.
//
// Note that because T is only used for a result parameter,
// function argument type inference does not work when calling
// this function.
func FromStrings[T Setter](s []string) []T {
	result := make([]T, len(s))
	for i, v := range s {
		result[i].Set(v)
	}
	return result
}
// Now let's see some calling code (this example is invalid).

// Settable is an integer type that can be set from a string.
type Settable int

// Set sets the value of *p from a string.
func (p *Settable) Set(s string) {
	i, _ := strconv.Atoi(s) // real code should not ignore the error
	*p = Settable(i)
}

func F() {
	// INVALID
	nums := FromStrings[Settable]([]string{"1", "2"})
	// Here we want nums to be []Settable{1, 2}.
	...
}

// So let's rewrite F to use *Settable instead.

func F() {
	// Compiles but does not work as desired.
	// This will panic at run time when calling the Set method.
	nums := FromStrings[*Settable]([]string{"1", "2"})
	...
}


// To repeat, we can‘t use Settable because it doesn’t have a Set method, and we can‘t use 
// *Settable because then we can’t create a slice of type Settable.

// What we can do is pass both types.

// Setter2 is a type constraint that requires that the type
// implement a Set method that sets the value from a string,
// and also requires that the type be a pointer to its type parameter.
type Setter2[B any] interface {
	Set(string)
	*B // non-interface type constraint element
}

// FromStrings2 takes a slice of strings and returns a slice of T,
// calling the Set method to set each returned value.
//
// We use two different type parameters so that we can return
// a slice of type T but call methods on *T aka PT.
// The Setter2 constraint ensures that PT is a pointer to T.
func FromStrings2[T any, PT Setter2[T]](s []string) []T {
	result := make([]T, len(s))
	for i, v := range s {
		// The type of &result[i] is *T which is in the type set
		// of Setter2, so we can convert it to PT.
		p := PT(&result[i])
		// PT has a Set method.
		p.Set(v)
	}
	return result
}


// We can then call FromStrings2 like this:

func F2() {
	// FromStrings2 takes two type parameters.
	// The second parameter must be a pointer to the first.
	// Settable is as above.
	nums := FromStrings2[Settable, *Settable]([]string{"1", "2"})
	// Now nums is []Settable{1, 2}.
	...
}

// This approach works as expected, but it is awkward to have to repeat Settable in the type 
// arguments. Fortunately, constraint type inference makes it less awkward. Using constraint 
// type inference we can write

func F3() {
	// Here we just pass one type argument.
	nums := FromStrings2[Settable]([]string{"1", "2"})
	// Now nums is []Settable{1, 2}.
	...
}

//There is no way to avoid passing the type argument Settable. But given that type argument, 
// constraint type inference can infer the type argument *Settable for the type parameter PT.

// As before, we create a mapping of known type arguments:

// {T -> Settable}

// We then unify each type parameter with a structural constraint. In this case, we unify PT 
// with the single type of Setter2[T], which is *T. The mapping is now

// {T -> Settable, PT -> *T}

// We then replace T with Settable throughout, giving us:

// {T -> Settable, PT -> *Settable}

// After this nothing changes, and we are done. Both type arguments are known.

// Index returns the index of e in s, or -1 if not found.
func Index[T Equaler](s []T, e T) int {
	for i, v := range s {
		if e.Equal(v) {
			return i
		}
	}
	return -1
}

// In order to write the Equaler constraint, we have to write a constraint that can refer 
// to the type argument being passed in. The easiest way to do this is to take advantage of 
// the fact that a constraint does not have to be a defined type, it can simply be an 
// interface type literal. This interface type literal can then refer to the type parameter.

// Index returns the index of e in s, or -1 if not found.
func Index[T interface { Equal(T) bool }](s []T, e T) int {
	// same as above
}

// This version of Index would be used with a type like equalInt defined here:

// equalInt is a version of int that implements Equaler.
type equalInt int

// The Equal method lets equalInt implement the Equaler constraint.
func (a equalInt) Equal(b equalInt) bool { return a == b }

// indexEqualInts returns the index of e in s, or -1 if not found.
func indexEqualInt(s []equalInt, e equalInt) int {
	// The type argument equalInt is shown here for clarity.
	// Function argument type inference would permit omitting it.
	return Index[equalInt](s, e)
}
// in this example, when we pass equalInt to Index, we check whether equalInt implements 
// the constraint interface { Equal(T) bool }. The constraint has a type parameter, so we 
// replace the type parameter with the type argument, which is equalInt itself. That gives 
// us interface { Equal(equalInt) bool }. The equalInt type has an Equal method with that 
// signature, so all is well, and the compilation succeeds.

// a constraint may use both constraint elements and methods.

// StringableSignedInteger is a type constraint that matches any
// type that is both 1) defined as a signed integer type;
// 2) has a String method.
type StringableSignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
	String() string
}

// he rules for type sets define what this means. The type set of the union element is the 
// set of all types whose underlying type is one of the predeclared signed integer types. The 
//type set of String() string is the set of all types that define that method. The type set 
// of StringableSignedInteger is the intersection of those two type sets. The result is the set 
// of all types whose underlying type is one of the predeclared signed integer types and that 
// defines the method String() string. 


// An example of a type argument that would be permitted is MyInt, defined as:

// MyInt is a stringable int.
type MyInt int

// The String method returns a string representation of mi.
func (mi MyInt) String() string {
	return fmt.Sprintf("MyInt(%d)", mi)
}

// a constraint element may be a type literal.

type byteseq interface {
	string | []byte
}
// The usual rules apply: the type argument for this constraint may be string or []byte; 
// a generic function with this constraint may use any operation permitted by both string and
// []byte.

// The byteseq constraint permits writing generic functions that work for either string or []byte types.

// Join concatenates the elements of its first argument to create a
// single value. sep is placed between elements in the result.
// Join works for string and []byte types.
func Join[T byteseq](a []T, sep T) (ret T) {
	if len(a) == 0 {
		// Use the result parameter as a zero value;
		// see discussion of zero value in the Issues section.
		return ret
	}
	if len(a) == 1 {
		// We know that a[0] is either a string or a []byte.
		// We can append either a string or a []byte to a []byte,
		// producing a []byte. We can convert that []byte to
		// either a []byte (a no-op conversion) or a string.
		return T(append([]byte(nil), a[0]...))
	}
	// We can call len on sep because we can call len
	// on both string and []byte.
	n := len(sep) * (len(a) - 1)
	for _, v := range a {
		// Another case where we call len on string or []byte.
		n += len(v)
	}

	b := make([]byte, n)
	// We can call copy to a []byte with an argument of
	// either string or []byte.
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	// As above, we can convert b to either []byte or string.
	return T(b)
}

// For composite types (string, pointer, array, slice, struct, function, map, channel) we 
// impose an additional restriction: an operation may only be used if the operator accepts 
// identical input types (if any) and produces identical result types for all of the types 
// in the type set. To be clear, this additional restriction is only imposed when a 
// composite type appears in a type set. It does not apply when a composite type is formed 
// from a type parameter outside of a type set, as in var v []T for some type parameter T.

// structField is a type constraint whose type set consists of some
// struct types that all have a field named x.
type structField interface {
	struct { a int; x int } |
		struct { b int; x float64 } |
		struct { c int; x uint64 }
}

// This function is INVALID.
func IncrementX[T structField](p *T) {
	v := p.x // INVALID: type of p.x is not the same for all types in set
	v++
	p.x = v
}

// sliceOrMap is a type constraint for a slice or a map.
type sliceOrMap interface {
	[]int | map[int]int
}

// Entry returns the i'th entry in a slice or the value of a map
// at key i. This is valid as the result of the operator is always int.
func Entry[T sliceOrMap](c T, i int) int {
	// This is either a slice index operation or a map key lookup.
	// Either way, the index and result types are type int.
	return c[i]
}

// sliceOrFloatMap is a type constraint for a slice or a map.
type sliceOrFloatMap interface {
	[]int | map[float64]int
}

// This function is INVALID.
// In this example the input type of the index operation is either
// int (for a slice) or float64 (for a map), so the operation is
// not permitted.
func FloatEntry[T sliceOrFloatMap](c T) int {
	return c[1.0] // INVALID: input type is either int or float64.
}

// Imposing this restriction makes it easier to reason about the type of some operation in 
// a generic function. It avoids introducing the notion of a value with a constructed type 
// set based on applying some operation to each element of a type set.


// A type literal in a constraint element can refer to type parameters of the constraint. 
// In this example, the generic function Map takes two type parameters. The first type 
// parameter is required to have an underlying type that is a slice of the second type 
// parameter. There are no constraints on the second type parameter.

// SliceConstraint is a type constraint that matches a slice of
// the type parameter.
type SliceConstraint[T any] interface {
	~[]T
}

// Map takes a slice of some element type and a transformation function,
// and returns a slice of the function applied to each element.
// Map returns a slice that is the same type as its slice argument,
// even if that is a defined type.
func Map[S SliceConstraint[E], E any](s S, f func(E) E) S {
	r := make(S, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// MySlice is a simple defined type.
type MySlice []int

// DoubleMySlice takes a value of type MySlice and returns a new
// MySlice value with each element doubled in value.
func DoubleMySlice(s MySlice) MySlice {
	// The type arguments listed explicitly here could be inferred.
	v := Map[MySlice, int](s, func(e int) int { return 2 * e })
	// Here v has type MySlice, not type []int.
	return v
}

// We showed other examples of this earlier in the discussion of constraint type inference.

// in a function with two type parameters From and To, a value of type From may be converted 
// to a value of type To if all the types in the type set of From's constraint can be 
// converted to all the types in the type set of To's constraint.

// This is a consequence of the general rule that a generic function may use any operation 
// that is permitted by all types listed in the type set.

// For example:

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func Convert[To, From integer](from From) To {
	to := To(from)
	if From(to) != from {
		panic("conversion out of range")
	}
	return to
}
// The type conversions in Convert are permitted because Go permits every integer type to be 
// converted to every other integer type.


// Some functions use untyped constants. An untyped constant is permitted with a value of a 
// type parameter if it is permitted with every type in the type set of the type parameter's 
// constraint.

// As with type conversions, this is a consequence of the general rule that a generic function 
// may use any operation that is permitted by all types in the type set.

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func Add10[T integer](s []T) {
	for i, v := range s {
		s[i] = v + 10 // OK: 10 can convert to any integer type
	}
}

// This function is INVALID.
func Add1024[T integer](s []T) {
	for i, v := range s {
		s[i] = v + 1024 // INVALID: 1024 not permitted by int8/uint8
	}
}

// When a constraint embeds another constraint, the type set of the outer constraint is the 
// intersection of all the type sets involved. If there are multiple embedded types, 
// intersection preserves the property that any type argument must satisfy the requirements 
// of all constraint elements.

// Addable is types that support the + operator.
type Addable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~complex64 | ~complex128 |
		~string
}

// Byteseq is a byte sequence: either string or []byte.
type Byteseq interface {
	~string | ~[]byte
}

// AddableByteseq is a byte sequence that supports +.
// This is every type that is both Addable and Byteseq.
// In other words, just the type set ~string.
type AddableByteseq interface {
	Addable
	Byteseq
}

// An embedded constraint may appear in a union element. The type set of the union is, as 
// usual, the union of the type sets of the elements listed in the union.

// Signed is a constraint with a type set of all signed integer
// types.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a constraint with a type set of all unsigned integer
// types.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer is a constraint with a type set of all integer types.
type Integer interface {
	Signed | Unsigned
}
