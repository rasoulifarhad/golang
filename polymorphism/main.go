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
	fmt.Printf("sending user email to  %s<%s>\n",
		u.name,
		u.email)
}

type admin struct {
	name  string
	email string
}

func (a admin) notify() {
	fmt.Printf("sending admin email to  %s<%s>\n",
		a.name,
		a.email)

}

func sendNotification(n notifier) {
	n.notify()

}

func main() {
	farhad := user{name: "farhad", email: "farhad@gooo.com"}
	sendNotification(&farhad)

	amir := admin{"amir", "amir@ff.com"}
	sendNotification(&amir)
}
