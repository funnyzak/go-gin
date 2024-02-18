package notification

import (
	"net/smtp"
	"strconv"
)

type SMTPPayload struct {
	Host      string
	Port      int
	Security  bool
	IgnoreTLS bool
	Username  string
	Password  string
	From      string
	To        string
	Cc        string
	Bcc       string
}

type SMTP struct {
	Payload SMTPPayload
}

func (s *SMTP) Send(title string, message string) error {
	auth := smtp.PlainAuth("", s.Payload.Username, s.Payload.Password, s.Payload.Host)
	to := []string{s.Payload.To}
	msg := []byte("To: " + s.Payload.To + "\r\n" +
		"Subject: " + title + "\r\n" +
		"\r\n" +
		message + "\r\n")
	err := smtp.SendMail(s.Payload.Host+":"+strconv.Itoa(s.Payload.Port), auth, s.Payload.From, to, msg)
	if err != nil {
		return err
	}
	return nil
}
