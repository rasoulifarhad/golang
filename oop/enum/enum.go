package enum

import "fmt"

type Size uint8

const (
	small Size = iota
	medium
	large
	extraLarge
)

func (s Size) toString() {
	switch s {
	case small:
		fmt.Println("Small")
	case medium:
		fmt.Println("Medium")
	case large:
		fmt.Println("Large")
	case extraLarge:
		fmt.Println("Extra Large")
	default:
		fmt.Println("Invalid Size entry")
	}
}

func main() {
	fmt.Println(small)
	fmt.Println(medium)
	fmt.Println(large)
	fmt.Println(extraLarge)
}
