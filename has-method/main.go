package main

import (
	"fmt"
	"io"
)

type A int
type B int

func (b B) M(x int) string {
	return fmt.Sprint(b, ": ", x)
}

func check(v interface{}) bool {
	_, has := v.(interface{ M(int) string })
	return has
}

func main() {
	var a A = 123
	var b B = 789
	fmt.Println(check(a)) // false
	fmt.Println(check(b)) // true

	//How to simulate for i in 0..N in some other languages?
	const N = 5

	for i := range [N]struct{}{} {
		fmt.Println(i)
	}
	for i := range [N][0]int{} {
		fmt.Println(i)
	}
	for i := range (*[N]int)(nil) {
		fmt.Println(i)
	}

	//How to detect native word size at compile time?
	//This tip is Go unrelated.

	const Is64bitArch = ^uint(0)>>63 == 1
	fmt.Println("Is64bitArch=", Is64bitArch)

	const Is32bitArch = ^uint(0)>>63 == 0
	fmt.Println("Is32bitArch=", Is32bitArch)
	const WordBits = 32 << (^uint(0) >> 63) // 64 or 32

}

type MyReader uint16

func NewMyReader() *MyReader {
	var mr MyReader
	return &mr
}

func (mr *MyReader) Read(data []byte) (int, error) {
	switch len(data) {
	default:
		*mr = MyReader(data[0])<<8 | MyReader(data[1])
		return 2, nil
	case 2:
		*mr = MyReader(data[0])<<8 | MyReader(data[1])
	case 1:
		*mr = MyReader(data[0])
	case 0:
	}
	return len(data), io.EOF
}

// Any of the following three lines ensures
// type *MyReader implements io.Reader.
var _ io.Reader = NewMyReader()
var _ io.Reader = (*MyReader)(nil)

func _() { _ = io.Reader(nil).(*MyReader) }
