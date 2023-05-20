// Sample program to show what happens when the outer and inner
// types implement the same interface.
package main

import (
	"fmt"
)

type notifier interface {
	notify()
}
type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	user
	level string
}

func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)

}
func sendNotif(n notifier) {
	n.notify()
}
func main() {

	ad := admin{
		user: user{
			name:  "farhad rasouli",
			email: "farhad@nnn",
		},
		level: "super",
	}

	sendNotif(&ad)

	ad.user.notify()

	ad.notify()

}
