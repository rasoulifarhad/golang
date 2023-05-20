package struct2

import (
	"encoding/json"
	"fmt"
	"io"
)

type Codec interface {
	Encode(w io.Writer, v interface{}) error
	Decode(r io.Reader, v interface{}) error
}

type jsonCodec struct{}

func (*jsonCodec) Encode(w io.Writer, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}
func (*jsonCodec) Decode(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

var JSON Codec = (*jsonCodec)(nil)

type encode func(w io.Writer, v interface{}) error
type decode func(r io.Reader, v interface{}) error

type customCodec struct {
	encode
	decode
}

func (cc *customCodec) Encode(w io.Writer, v interface{}) error { return cc.encode(w, v) }
func (cc *customCodec) Decode(r io.Reader, v interface{}) error { return cc.decode(r, v) }

func NewCodec(enc encode, dec decode) *customCodec {
	return &customCodec{enc, dec}
}

// func main() {
// 	obj := "string"

// 	_ = JSON.Encode(os.Stdout, obj)

// 	json := NewCodec(
// 		func(w io.Writer, v interface{}) error { return json.NewEncoder(w).Encode(v) },
// 		func(r io.Reader, v interface{}) error { return json.NewDecoder(r).Decode(v) },
// 	)
// 	_ = json.Encode(os.Stdout, obj)
// }

func hello(id int, quit chan struct{}) {
	for {
		select {
		case <-quit:
			fmt.Printf("exited --> %d\n", id)
			return
		default:
			//fmt.Printf("hello --> %d\n", id)
		}
	}

}

//   func main() {
// 	quit := make(chan struct{})
// 	go hello(quit)
// 	// print hello for 10 ms
// 	time.Sleep(10 * time.Millisecond)
// 	quit <- struct{}{} // quit printing hello
//   }
