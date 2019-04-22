package main

import (
	"log"

	"github.com/LIYINGZHEN/gosmtp/pkg/config"
	"github.com/LIYINGZHEN/gosmtp/pkg/smtp"
)

func main() {
	mail := smtp.NewMail()

	config := config.New()
	gosmtp := smtp.New(config.Host, config.Port, config.Sender, config.Password)

	err := gosmtp.Send(mail)
	if err != nil {
		log.Printf("[Error] Unable to send emails: %v", err)
	}
}
