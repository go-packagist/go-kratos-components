package smtp

import (
	"net/smtp"

	"github.com/go-kratos-ecosystem/components/v2/x/mailer"
)

type smtpTransport struct {
	auth smtp.Auth
}

type Option func(*smtpTransport)

func Auth(auth smtp.Auth) Option {
	return func(t *smtpTransport) {
		t.auth = auth
	}
}

func NewTransport(opts ...Option) mailer.Transport {
	s := &smtpTransport{}

	for _, o := range opts {
		o(s)
	}

	return s
}

func (t *smtpTransport) Send(message mailer.Message) error {
	return nil
	// return smtp.SendMail(message.FromAddress(), t.auth, message.From, message.To, message.Body)
}
