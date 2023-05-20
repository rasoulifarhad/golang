package overloading

import "fmt"

// - Different number of parameters but of the same type:
func add(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum

}

// - Different number of parameters and of different types:
// This case can be handled using both variadic function and empty interface
func handle(params ...interface{}) {
	fmt.Println("Handle func called with parameters:")
	for _, param := range params {
		fmt.Printf("%v\n", param)
	}

}

type person struct {
	name   string
	gender string
	age    int
}

func addPerson(args ...interface{}) error {
	if len(args) > 3 {
		return fmt.Errorf("wrong number of args passed")
	}

	p := person{}
	//0 is name
	//1 is gender
	//2 is age
	for i, arg := range args {
		switch i {
		case 0: // name
			name, ok := arg.(string)
			if !ok {
				return fmt.Errorf("name is not passed as string")
			}
			p.name = name
		case 1: // gender
			gender, ok := arg.(string)
			if !ok {
				return fmt.Errorf("gender is not passed as string")
			}
			p.gender = gender
		case 2: // age
			age, ok := arg.(int)
			if !ok {
				return fmt.Errorf("age is not passed as int")
			}
			p.age = age
		}
	}
	fmt.Printf("person struct: %v\n", p)
	return nil

}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(1, 2, 3))
	fmt.Println(add(1, 2, 3, 4))

	handle(1, "abc")
	handle("abc", "xyz", 3)
	handle(1, 2, 3, 4)

	err := addPerson("Tina", "Female", 20)
	if err != nil {
		fmt.Println("PersonAdd Error: " + err.Error())
	}

	err = addPerson("John", "Male")
	if err != nil {
		fmt.Println("PersonAdd Error: " + err.Error())
	}

	err = addPerson("Wick", 2, 3)
	if err != nil {
		fmt.Println("PersonAdd Error: " + err.Error())
	}
}
