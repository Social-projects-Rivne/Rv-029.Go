package mail

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
)

type mailer interface {
	Send(address []mail.Address, subject string, msg string) error
}

type SmtpMailerConfig struct {
	Connection struct {
		Host     string
		Port     int
		Username string
		Password string
		Auth     string
		Tls      bool
	}
	auth   smtp.Auth
	Sender struct {
		Name  string
		Email string
	}
}

type SmtpMailer struct {
	*SmtpMailerConfig
}

var Mailer *SmtpMailer

func InitFromConfig(config *SmtpMailerConfig) *SmtpMailer {
	mailer := &SmtpMailer{config}

	switch mailer.Connection.Auth {
	//case "login": //TODO:
	case "cram_md5":
		mailer.auth = smtp.CRAMMD5Auth(mailer.Connection.Username, mailer.Connection.Password)
	default:
		mailer.auth = smtp.PlainAuth("", mailer.Connection.Username, mailer.Connection.Password, mailer.Connection.Host)
	}

	return mailer
}

func (m *SmtpMailer) Send(toEmail, toName, subject, msg string) error {
	from := mail.Address{m.Sender.Name, m.Sender.Email}
	to := mail.Address{m.Sender.Name, m.Sender.Email}

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	headers["MIME-version"] = " 1.0;"
	headers["Content-Type"] = "text/html; charset=\"UTF-8\";"

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + msg

	// Connect to the SMTP Server
	c, err := smtp.Dial(fmt.Sprintf("%s:%d", m.Connection.Host, m.Connection.Port))
	if err != nil {
		log.Panic(err)
	} else {
		defer c.Close()
	}

	if m.Connection.Tls {
		// TLS config
		tlsconfig := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         m.Connection.Host,
		}

		c.StartTLS(tlsconfig)
	}

	// Auth
	if err = c.Auth(m.auth); err != nil {
		log.Printf("Error in utils/mail/mail.go error: %+v",err)
		return err
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Printf("Error in utils/mail/mail.go error: %+v",err)
		return err
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Printf("Error in utils/mail/mail.go error: %+v",err)
		return err
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Printf("Error in utils/mail/mail.go error: %+v",err)
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Printf("Error in utils/mail/mail.go error: %+v",err)
		return err
	}

	err = w.Close()
	if err != nil {
		log.Printf("Error in utils/mail/mail.go error: %+v",err)
		return err
	}

	c.Quit()

	return nil
}
