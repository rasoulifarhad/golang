package nonstructinterface

import (
	"fmt"
	"testing"
)

func TestDo(t *testing.T) {
	var a animal

	a = cat("smokey")
	a.breathe()
	a.walk()
}

func TestInterfaceEquality(t *testing.T) {
	var a animal
	var b animal
	var c animal
	var d animal
	var e animal

	a = lion{age: 10}
	b = lion{age: 10}
	c = lion{age: 5}

	if a == b {
		fmt.Println("a and b are equal")
	} else {
		fmt.Println("a and b are not equal")
	}

	if a == c {
		fmt.Println("a and c are equal")
	} else {
		fmt.Println("a and c are not equal")
	}

	if d == e {
		fmt.Println("d and e are equal")
	} else {
		fmt.Println("d and e are not equal")
	}
}
