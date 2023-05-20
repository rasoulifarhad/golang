package pointer

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestStruct(t *testing.T) {
	content := []byte(`{"foo":"bar"}`)

	var result1, result2 *Result
	//result2 = &Result{}
	err := json.Unmarshal(content, &result1)
	fmt.Println(result1, err)

	fmt.Println(result2)
	fmt.Println(&result2)

	err = json.Unmarshal(content, result2)
	fmt.Println(result2, err)

}

func TestStruct2(t *testing.T) {
	content := []byte(`{"foo":"bar"}`)

	var result1 *Result
	//result2 = &Result{}
	err := json.Unmarshal(content, &result1)

	fmt.Printf("%#v %v", result1, err) // &pointer.Result{Foo:"bar"} <nil>
}
func TestStruct3(t *testing.T) {
	content := []byte(`{"foo":"bar"}`)

	var result1 *PtResult
	//result2 = &Result{}
	err := json.Unmarshal(content, &result1)
	fmt.Printf("%#v %v", result1, err) // &pointer.PtResult{Foo:(*string)(0xc000052630)} <nil>

}

func TestInt(t *testing.T) {

	content := []byte("1")

	var result *int

	err := json.Unmarshal(content, &result)

	fmt.Println(*result, err)
}
