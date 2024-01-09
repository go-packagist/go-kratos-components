package mailer

type Mailer interface {
	Send(message Message) error
}

type mailer struct {
	transport Transport
}

func NewMailer(transport Transport) Mailer {
	return &mailer{
		transport: transport,
	}
}

func (m *mailer) Send(message Message) error {
	return m.transport.Send(message)
}
