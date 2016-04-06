package emim

import "github.com/scorredoira/email"

type Message struct {
	*email.Message
}

func Create(m *email.Message) *Message {
	return &Message{
		m,
	}
}

func (m *Message) From() string {
	return m.Message.From
}

func (m *Message) To() []string {
	return m.Tolist()
}
