package embeddingInterface

import (
	"fmt"
	"testing"
)

func TestDo(t *testing.T) {
	d := dog{age: 5}
	p1 := pet1{name: "Milo", a: d}

	fmt.Println(p1.name)
	// p1.breathe()
	// p1.walk()
	p1.a.breathe()
	p1.a.walk()

	p2 := pet2{name: "Oscar", animal: d}
	fmt.Println(p1.name)
	p2.breathe()
	p2.walk()
	p2.animal.breathe()
	p2.animal.walk()
}

func TestDo3(t *testing.T) {
	d := dog{age: 5}
	p1 := pet1{name: "Milo", a: d}

	fmt.Println(p1.name)
	//p1.breathe()
	//p1.walk()
	p1.a.breathe()
	p1.a.walk()

}
