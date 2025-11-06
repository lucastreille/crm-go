package notification

import "fmt"

type Notifier interface {
	Send(message string) error
}

type EmailNotifier struct{}

func (e EmailNotifier) Send(message string) error {
	fmt.Println("📧 Email envoyé :", message)
	return nil
}

type SmsNotifier struct{}

func (s SmsNotifier) Send(message string) error {
	fmt.Println("📱 SMS envoyé :", message)
	return nil
}

func NotifyAll(notifiers []Notifier, message string) {
	for _, n := range notifiers {
		n.Send(message)
	}
}
