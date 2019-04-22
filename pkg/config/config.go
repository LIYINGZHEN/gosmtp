package config

import (
	"os"
)

// Config contains all the environment varialbes.
type Config struct {
	Host     string
	Port     string
	Password string
	Sender   string
}

func New() Config {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "465"
	}

	host, ok := os.LookupEnv("HOST")
	if !ok {
		host = "smtp.gmail.com"
	}

	sender, ok := os.LookupEnv("SENDER")
	if !ok {
		sender = "account@gmail.com"
	}

	password, ok := os.LookupEnv("PASSWORD")
	if !ok {
		password = "password"
	}

	return Config{
		Port:     port,
		Host:     host,
		Password: password,
		Sender:   sender,
	}
}
