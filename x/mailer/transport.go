package mailer

type Transport interface {
	Send(message Message) error
}
