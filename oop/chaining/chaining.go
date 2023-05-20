//      Method Chaining in Go
// For method chaining to be possible, the methods in the chain should
// return the receiver. Returning the receiver for the last method in
// the chain is optional.
package chaining

import "fmt"

type employee struct {
	name   string
	age    int
	salary int
}

func (e employee) printName() employee {
	fmt.Printf("name: %s\n", e.name)
	return e
}

func (e employee) printAge() employee {
	fmt.Printf("age: %d\n", e.age)
	return e
}

func (e employee) printSalary() {
	fmt.Printf("salary: %d\n", e.salary)

}

func main() {
	emp := employee{name: "Sam", age: 31, salary: 2000}
	emp.printName().printAge().printSalary()
}
