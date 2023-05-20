// Go supports inheritance by embedding struct or using interface. There are different ways
// of doing it and each having some limitations. The different ways are:

//    1- By using embedded struct – The parent struct is embedded in child struct. The
//       limitation is that subtyping is not possible with this approach. You cannot pass
//       the child struct to a function that expects base.
//
//    2- By using interfaces – Subtyping is possible but the limitation is that one has no
//       way to refer to common properties. Refer this link for more details

//    3- By using interface + struct – This fixes the limitations of above two approach but
//       one limitation is that overriding methods is not possible. But there is workaround.

package main

import (
	"fmt"
)

type iBase interface {
	say()
}

type base struct {
	color string
	clear func()
}

func (b *base) say() {
	//fmt.Println(b.color)
	//b.clear()
	b.clear()
}

// func (b *base) clear() {
// 	fmt.Println("Clear from base's function")
// }

type child struct {
	base  //embedding
	style string
}

// func (b *child) clear() {
// 	fmt.Println("Clear from child's function")
// }

func check(b iBase) {
	b.say()
}

func main() {
	//base := base{color: "red"}

	//child := &child{
	//	base:  base,
	//	style: "somestyle",
	//}
	//child.say()
	//check(child)

	base := base{
		color: "red",
		clear: func() {
			fmt.Println("Clear from child's functio")
		}}
	child := &child{
		base:  base,
		style: "somStyle",
	}
	child.say()
	//e := employee.New("farhad", "rasouli", 30, 20)

	//e.LeavesRemaining()

}
