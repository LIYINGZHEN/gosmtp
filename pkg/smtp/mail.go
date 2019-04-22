package smtp

import (
	"bufio"
	"bytes"
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

func NewMail(sender string) *Mail {
	var (
		to      string
		subject string
	)

	fmt.Println("Please enter receivers. (Use comma as a Separator) Example: a@gmail.com,b@gmail.com")
	fmt.Scan(&to)
	receivers := strings.Split(to, ",")

	fmt.Println("Please enter the email subject.")
	fmt.Scan(&subject)

	fmt.Println("Please enter the message. You can end with '|'")
	scanner := bufio.NewScanner(os.Stdin)
	var body string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "|" {
			break
		}
		body += line + "\n"
	}

	return &Mail{
		Sender:  sender,
		To:      receivers,
		Subject: subject,
		Body:    body,
	}
}

func (mail *Mail) buildMessage() string {
	var buf bytes.Buffer
	buf.WriteString("From: " + mail.Sender)
	buf.WriteString("\r\n")
	if len(mail.To) > 0 {
		buf.WriteString("To: " + strings.Join(mail.To, ";"))
		buf.WriteString("\r\n")
	}
	buf.WriteString("Subject: " + mail.Subject)
	buf.WriteString("\r\n")
	buf.WriteString(mail.Body)

	return buf.String()
}
