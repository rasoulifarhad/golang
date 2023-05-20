// Sample program to show how embedded types work with interfaces.
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

	ad.user.notify()

	ad.notify()

	sendNotif(&ad)
}
