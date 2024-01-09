package mailer

import "net/textproto"

type Attachment struct {
}

type Message struct {
	replyTo     []string
	from        string
	to          []string
	bcc         []string
	cc          []string
	subject     string
	text        []byte // Plaintext message (optional)
	html        []byte // Html message (optional)
	sender      string // override From as SMTP envelope sender (optional)
	headers     textproto.MIMEHeader
	attachments []*Attachment
	readReceipt []string
}

type MessageOption func(*Message)

func NewMessage(opts ...MessageOption) *Message {
	m := &Message{}

	for _, o := range opts {
		o(m)
	}

	return m
}

func (m *Message) ReplyTo(replyTo ...string) *Message {
	m.replyTo = replyTo
	return m
}

func (m *Message) From(from string) *Message {
	m.from = from
	return m
}

func (m *Message) To(to ...string) *Message {
	m.to = to
	return m
}

func (m *Message) Cc(cc ...string) *Message {
	m.cc = cc
	return m
}

func (m *Message) Bcc(bcc ...string) *Message {
	m.bcc = bcc
	return m
}

func (m *Message) Subject(subject string) *Message {
	m.subject = subject
	return m
}

func (m *Message) Text(text []byte) *Message {
	m.text = text
	return m
}

func (m *Message) Html(html []byte) *Message {
	m.html = html
	return m
}

func (m *Message) Sender(sender string) *Message {
	m.sender = sender
	return m
}

func (m *Message) Attach(attachments ...*Attachment) *Message {
	m.attachments = append(m.attachments, attachments...)
	return m
}
