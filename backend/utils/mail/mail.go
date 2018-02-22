package mail

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/mail"
	"net/smtp"
	"path/filepath"
)

type mailer interface {
	Send(address []mail.Address, subject string, msg string) error
}

type smtpMailer struct {
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

var Mailer *smtpMailer

func (m *smtpMailer) Send(toEmail, toName, subject, msg string) error {

	filename, _ := filepath.Abs("./backend/config/mail.yml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	Mailer = &smtpMailer{}

	err = yaml.Unmarshal(yamlFile, &Mailer)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	switch Mailer.Connection.Auth {
	//case "login": //TODO:
	case "cram_md5":
		Mailer.auth = smtp.CRAMMD5Auth(Mailer.Connection.Username, Mailer.Connection.Password)
	default:
		Mailer.auth = smtp.PlainAuth("", Mailer.Connection.Username, Mailer.Connection.Password, Mailer.Connection.Host)
	}

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
		return err
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		return err
	}

	if err = c.Rcpt(to.Address); err != nil {
		return err
	}

	// Data
	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	c.Quit()

	return nil
}
