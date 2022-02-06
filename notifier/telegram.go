package notifier

type Telegram struct {
	Token   string
	Message string
	INotifier
}

func (t *Telegram) SetMessage(msg string) {
	t.Message = msg
}

func (t *Telegram) SendMessage(recipient string) error {

	return nil
}
