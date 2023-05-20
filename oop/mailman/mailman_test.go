package mailman

import (
	"fmt"
	"net/mail"
	"testing"
)

type testEmailSender struct {
	lastSubject string
	lastBody    string
	lastTo      []*mail.Address
}

func (t *testEmailSender) Send(subject string, body string, to ...*mail.Address) {
	t.lastSubject = subject
	t.lastBody = body
	t.lastTo = to
	fmt.Printf("test mail sended.\nsubject: %s\nbody: %s\n", subject, body)

}

var eSender = (*testEmailSender)(nil)

func TestSendWelcomEmail(t *testing.T) {
	sender := &testEmailSender{}
	to1 := &mail.Address{Name: "farhad"}
	to2 := &mail.Address{Name: "farhad2"}
	SendWelcomMail(sender, to1, to2)
}
