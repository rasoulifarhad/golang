package wrapper

import (
	"compress/zlib"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestReflect(t *testing.T) {
	callme(hello, 1)

	callme(hiya, "buddy")

	callme(awesome, 5, "buddy")
}

func TestSumWrapper(t *testing.T) {
	printWrapper(3, 4, sum)
}

func TestMulWrapper(t *testing.T) {
	printWrapper(2, 4, multiply)
}

func TestClouser(t *testing.T) {
	msg := "variable from outer scope"
	s := func(x, y int) int {
		fmt.Println(msg)
		return x - y
	}
	printWrapper(5, 2, s)
}

func TestLogging(t *testing.T) {
	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))

	http.ListenAndServe(":8080", nil)
}

func TestIo(t *testing.T) {
	src, err := os.Open("/home/farhad/tippppp")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	dest, err := os.Create("/home/farhad/tippppp-new")

	if err != nil {
		log.Fatal(err)
	}
	defer dest.Close()

	if _, err := io.Copy(dest, src); err != nil {
		log.Fatal(err)
	}
}

func TestIoZip(t *testing.T) {
	src, err := os.Open("/home/farhad/tippppp")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	dest, err := os.Create("/home/farhad/tippppp-new2")

	if err != nil {
		log.Fatal(err)
	}
	defer dest.Close()
	zdest := zlib.NewWriter(dest)
	defer zdest.Close()

	if _, err := io.Copy(zdest, src); err != nil {
		log.Fatal(err)
	}
}

func TestDefer1(t *testing.T) {
	log.Println("open database connection")
	defer log.Println("close database")

	log.Println("open file")
	defer log.Println("close file")

	log.Fatal("stopped")

}

func TestDefer(t *testing.T) {
	var fatalErr error
	defer func() {
		if fatalErr != nil {
			log.Fatalln(fatalErr)
		}
	}()

	log.Println("open database connection")
	defer log.Println("close database")

	log.Println("open file")
	defer log.Println("close file")

	// simluate error
	fatalErr = errors.New("soemthing went wrong")
	return
}
