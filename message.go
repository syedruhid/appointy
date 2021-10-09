package main

import (
	"github.com/go-mail/mail"
	"regexp"
	"strings"
)

var rxEmail = regexp.MustCompile(".+@.+\\..+")

type Message struct {
	email    string
	password string
	Errors   map[string]string
}

func (msg *Message) Validate() bool {
	msg.Errors = make(map[string]string)

	match := rxEmail.Match([]byte(msg.email))
	if match == false {
		msg.Errors["email"] = "Please enter a valid Username"
	}

	if strings.TrimSpace(msg.password) == "" {
		msg.Errors["password"] = "Please enter your Password"
	}

	return len(msg.Errors) == 0
}

func (msg *Message) Deliver() error {
	email := mail.NewMessage()
	email.SetHeader("To", "admin@example.com")
	email.SetHeader("From", "server@example.com")
	email.SetHeader("Reply-To", msg.email)
	email.SetHeader("Subject", "New message via Contact Form")
	email.SetBody("text/plain", msg.password)

	username := "your_username"
	password := "your_password"

	return mail.NewDialer("smtp.mailtrap.io", 25, username, password).DialAndSend(email)
}
