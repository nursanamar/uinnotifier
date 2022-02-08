package notifier

import (
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

func TestSendMail(t *testing.T) {

	errenv := godotenv.Load("../.env")
	if errenv != nil {
		t.Fatal("Error loading .env file")
	}

	smtpPort, err := strconv.ParseInt(os.Getenv("SMTP_PORT"), 10, 32)
	if err != nil {
		smtpPort = 465
	}

	mail := &Mail{
		host:     os.Getenv("SMTP_HOST"),
		port:     int(smtpPort),
		user:     os.Getenv("SMTP_USER"),
		password: os.Getenv("SMTP_PASSWORD"),
	}

	mail.SetMessage("Halo")
	err = mail.SendMessage("nursan011@gmail.com")

	if err != nil {
		t.Fail()
	}
}
