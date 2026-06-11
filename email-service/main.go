package main

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	from := os.Getenv("EMAIL_ID")
	password := os.Getenv("PASSWORD")

	to := []string{
		os.Getenv("RECEIVERS_MAIL_ID"),
	}

	message := []byte(
		"Subject: Greetings from golang email service\r\n" +
			"\r\n" +
			"Hello from Go email service!",
	)

	auth := smtp.PlainAuth(
		"",
		from,
		password,
		"smtp.gmail.com",
	)

	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		from,
		to,
		message,
	)

	if err != nil {
		fmt.Println("Failed:", err)
		return
	}

	fmt.Println("Email Sent!")
}
