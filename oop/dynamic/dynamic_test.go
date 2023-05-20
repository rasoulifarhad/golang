package dynamic

import (
	"fmt"
	"io"
	"os"
	"syscall"
	"testing"
	"time"
	"unsafe"
)

func TestOpenFile(t *testing.T) {
	_, err := os.Open("nptFount.txt")

	if err != nil {
		t.Fatal(err)
	}
}

func TestOpenFileBoom(t *testing.T) {
	f, err := os.Open("nptFount.txt")

	if err != nil {
		t.Error(err)
	}
	f.Close()
}

// DRY it up, of course, by moving the repetitive assertion logic to a helper:
func TestOpenFileDRY(t *testing.T) {
	_, err := os.Open("notFound.txt")

	check(t, err)
}

func check(t *testing.T, err error) {
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
}

func TestNotFountPanic(t *testing.T) {
	_, err := os.Open("notFound.txt")
	checkPanic(err)
}

func checkPanic(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func TestUnitSize(t *testing.T) {
	const uintSize = 32 << (^uint(0) >> 32 & 1)

	fmt.Println(uintSize)
}

func TestConstantValuType(t *testing.T) {

	const timeout = 500 * time.Millisecond

	fmt.Println("the timeout is : ", timeout)
	fmt.Printf("timeout: type %T , value %v , value+ %+v\n", timeout, timeout, timeout)
}

// var (
// 	Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
// 	Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
// 	Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
// )
//their type is *os.File not the respective io.Reader or io.Writer interfaces.
//Could we change the definition of os.Stdout and friends so that they retain the observable
//behaviour of reading and writing, but remain immutable? It turns out, we can do this easily
//with constants.
func TestFileDescriptor(t *testing.T) {
	fmt.Fprintf(os.Stdout, "Hello dotGo\n")
	syscall.Write(1, []byte("Hello dotGo\n"))
}

func TestStd(t *testing.T) {

	fmt.Fprintf(Stdout, "Hello ")
}

func TestChangeEOFError(t *testing.T) {
	fmt.Printf("EOf: %T , %v\n ", io.EOF, io.EOF)
	//io.EOF = nil

	fmt.Printf("EOf: %T , %v\n ", io.EOF, io.EOF)

	fmt.Printf("nil: %T , %v\n ", nil, nil)

	fmt.Printf("myEOF: %T , %v\n ", myEOF, myEOF)

	fmt.Printf("myEOF == io.EOF : %v\n", myEOF == io.EOF)

	const err1 = Error("EOF")
	const err2 = Error("EOF")
	fmt.Printf("err1: %T , %v\n ", err1, err1)
	fmt.Printf("err2: %T , %v\n ", err2, err2)
	fmt.Printf("err1 == err2 : %v\n", err1 == err2)

}

var m map[int]int
var p uintptr

func TestMap(t *testing.T) {
	fmt.Println(unsafe.Sizeof(m), unsafe.Sizeof(p)) // 8 8 (linux/amd64)
}

func TestTypeValue(t *testing.T) {
	fmt.Printf("Type: %T  Value:%+v\n", w, w)
}
