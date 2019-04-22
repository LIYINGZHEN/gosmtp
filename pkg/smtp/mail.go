package smtp

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Mail struct {
	Sender  string
	To      []string
	Subject string
	Body    string
}

func NewMail() *Mail {
	var (
		sender  string
		to      string
		subject string
	)

	fmt.Println("input user as:user@xx.xx")
	fmt.Scan(&sender)

	fmt.Println("input sent list as:to1@xx.xx,to2@xx.xx,...")
	fmt.Scan(&to)
	tolist := strings.Split(to, ",")

	fmt.Println("input subject")
	fmt.Scan(&subject)

	fmt.Println("input message and end with '|' then newline")
	inputReader := bufio.NewReader(os.Stdin)
	body, err := inputReader.ReadString('|')
	if err != nil {
		fmt.Println(err)
	}

	return &Mail{
		Sender:  sender,
		To:      tolist,
		Subject: subject,
		Body:    body,
	}
}

func (mail *Mail) BuildMessage() string {
	header := ""
	header += fmt.Sprintf("From: %s\r\n", mail.Sender)
	if len(mail.To) > 0 {
		header += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	}

	header += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	header += "\r\n" + mail.Body

	return header
}
