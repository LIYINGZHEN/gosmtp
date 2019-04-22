package main

import (
	"log"

	"github.com/LIYINGZHEN/gosmtp/pkg/config"
	"github.com/LIYINGZHEN/gosmtp/pkg/smtp"
)

func main() {
	config := config.New()
	gosmtp := smtp.New(config.Host, config.Port, config.Sender, config.Password)
	mail := smtp.NewMail(config.Sender)
	err := gosmtp.Send(mail)
	if err != nil {
		log.Printf("[Error] Unable to send emails: %v", err)
	}
}
