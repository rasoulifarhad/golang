package chaining

import "testing"

func TestChaining(t *testing.T) {
	emp := employee{name: "Sam", age: 31, salary: 2000}
	emp.printName().printAge().printSalary()
}
