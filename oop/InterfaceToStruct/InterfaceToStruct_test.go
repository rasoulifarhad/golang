package InterfaceToStruct

import (
	"testing"
)

func TestDo(t *testing.T) {
	newCustomer := NewCustomerEvent{Name: "x", Phone: "082213909101", Email: "xyz@gmail.com"}
	convert(newCustomer)
}
