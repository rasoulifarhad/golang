package wrapper

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
)

func callme(fn interface{}, params ...interface{}) (result []reflect.Value) {
	f := reflect.ValueOf(fn)
	if f.Type().NumIn() != len(params) {
		panic("incorrect number of parameters!")
	}
	inputs := make([]reflect.Value, len(params))

	for k, param := range params {
		inputs[k] = reflect.ValueOf(param)
	}
	return f.Call(inputs)

}

func hello(i int) {
	fmt.Println("hello " + strconv.Itoa(i))
}

func hiya(name string) {
	fmt.Println("hiya " + name)
}

func awesome(i int, name string) {
	fmt.Println("high " + strconv.Itoa(i) + ", " + name)
}

func sum(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func printWrapper(a, b int, fn func(int, int) int) {
	fmt.Printf("inputs: %d %d\n", a, b)
	res := fn(a, b)
	fmt.Printf("result: %d\n", res)

}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		f(w, r)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")
}
