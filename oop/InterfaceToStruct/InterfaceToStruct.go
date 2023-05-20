// We come around a situation sometimes in programming where an empty interface might a
// struct internally and we have to get the concrete struct out of it.
// For conversion of interface{} to a struct, we will use the library –
// https://github.com/mitchellh/mapstructure
package InterfaceToStruct

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type NewCustomerEvent struct {
	Name  string
	Phone string
	Email string
}

func convert(event interface{}) {
	c := NewCustomerEvent{}
	mapstructure.Decode(event, &c)
	fmt.Printf("Event is: %v", c)
}

func main() {
	newCustomer := NewCustomerEvent{Name: "x", Phone: "082213909101", Email: "xyz@gmail.com"}
	convert(newCustomer)
}
