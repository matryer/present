package main

import (
	"errors"
	"fmt"
	"testing"
)

// START OMIT

type MailSender interface {
	Send(to, subject, body string) error
	SendFrom(from, to, subject, body string) error
}

type MockSender struct {
	SendFunc     func(to, subject, body string) error
	SendFromFunc func(from, to, subject, body string) error
}

func (m MockSender) Send(to, subject, body string) error {
	return m.SendFunc(to, subject, body)
}
func (m MockSender) SendFrom(from, to, subject, body string) error {
	return m.SendFunc(from, to, subject, body)
}

// END OMIT

// TWO OMIT

func TestWelcomeEmail(t *testing.T) {
	errTest := errors.New("nope")
	var msg string

	sender := MockSender{
		SendFunc: func(to, subject, body string) error {
			msg = fmt.Sprintf("(%s) %s: %s", to, subject, body)
			return nil
		},
		SendFromFunc: func(from, to, subject, body string) error {
			return errTest
		},
	}

	SendWelcomeEmail(sender, "to", "subject", "body")

	if msg != "(to) subject: body" {
		t.Error("SendWelcomeEmail:", msg)
	}
}

// TWO END OMIT

func SendWelcomeEmail(sender MailSender, to, subject, body string) {

}
