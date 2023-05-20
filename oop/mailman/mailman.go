package mailman

import (
	"fmt"
	"net/mail"
)

type EmailSender interface {
	Send(subject string, body string, to ...*mail.Address)
}
type Mailman struct{}

func (m *Mailman) Send(subject string, body string, to ...*mail.Address) {
	// som code
	fmt.Printf("real mail sended.\nsubject: %s\nbody: %s\n", subject, body)
}

func New() *Mailman {
	return &Mailman{}
}

// func SendWelcomMail(m *Mailman, to ...*mail.Address) {
// 	m.Send("welcom mail", "welcom  boy", to...)
// }

func SendWelcomMail(m EmailSender, to ...*mail.Address) {
	SendMail(m, "welcom mail", "welcom  boy", to...)
}
func SendMail(m EmailSender, subject string, body string, to ...*mail.Address) {
	m.Send(subject, body, to...)
}
