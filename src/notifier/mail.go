package notifier

import (
	"gopkg.in/gomail.v2"
)

type Mail struct {
	host     string
	port     int
	user     string
	password string
	message  gomail.Message
	INotifier
}

func (m *Mail) SetHost(host string) {
	m.host = host
}

func (m *Mail) SetPort(port int) {
	m.port = port
}

func (m *Mail) SetUser(user string) {
	m.user = user
}

func (m *Mail) SetPassword(password string) {
	m.password = password
}

func (m *Mail) SetMessage(message string) {
	m.message = *gomail.NewMessage()

	m.message.SetHeader("From", m.user)
	m.message.SetHeader("Subject", "Notifikasi")
	m.message.SetBody("text/html", message)
}

func (m *Mail) SendMessage(recipient string) error {
	dialer := gomail.NewDialer(m.host, m.port, m.user, m.password)

	m.message.SetHeader("To", recipient)

	err := dialer.DialAndSend(&m.message)

	return err
}
