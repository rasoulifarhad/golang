// Package entities contains support for types of
// people in the system.
package entities

//type User struct {
//Name string
//email string
//}

type user struct {
	Name  string
	Email string
}

type Admin struct {
	user
	Rights int
}
