package notifier

import (
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type INotifier interface {
	SetMessage(msg string)
	SendMessage(recipient string) error
}

func MakeNotifier(name string) INotifier {
	smtpPort, err := strconv.ParseInt(os.Getenv("SMTP_PORT"), 10, 32)
	if err != nil {
		smtpPort = 465
	}

	return &Mail{
		host:      os.Getenv("SMTP_HOST"),
		port:      int(smtpPort),
		user:      os.Getenv("SMTP_USER"),
		password:  os.Getenv("SMTP_PASSWORD"),
		message:   gomail.Message{},
		INotifier: nil,
	}
}
