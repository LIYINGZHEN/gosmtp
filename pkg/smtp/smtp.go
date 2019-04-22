package smtp

import (
	"crypto/tls"
	"log"
	"net/smtp"
)

type Smtp struct {
	Host      string
	Port      string
	Email     string
	Password  string
	TlsConfig *tls.Config
	Auth      smtp.Auth
}

func New(host string, port string, email string, password string) *Smtp {
	return &Smtp{
		Host:     host,
		Port:     port,
		Email:    email,
		Password: password,
		TlsConfig: &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         host,
		},
		Auth: smtp.PlainAuth("", email, password, host),
	}
}

func (s *Smtp) serverName() string {
	return s.Host + ":" + s.Port
}

func (s *Smtp) Send(mail *Mail) error {
	conn, err := tls.Dial("tcp", s.serverName(), s.TlsConfig)
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(conn, s.Host)
	if err != nil {
		return err
	}

	if err = client.Auth(s.Auth); err != nil {
		return err
	}

	if err = client.Mail(mail.Sender); err != nil {
		return err
	}

	for _, k := range mail.To {
		log.Println("sending to: ", k)
		if err = client.Rcpt(k); err != nil {
			return err
		}
	}

	w, err := client.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(mail.buildMessage()))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	client.Quit()

	return nil
}
