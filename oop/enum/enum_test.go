package enum

import (
	"fmt"
	"testing"
)

func TestDo(t *testing.T) {
	fmt.Println(small)
	fmt.Println(medium)
	fmt.Println(large)
	fmt.Println(extraLarge)
}

func TestToString(t *testing.T) {
	var m Size = 1
	m.toString()
}
