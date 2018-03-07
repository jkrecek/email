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

func (m *Message) GetFrom() string {
	return m.Message.From.Address
}

func (m *Message) GetTo() []string {
	return m.Tolist()
}

func (m *Message) GetBytes() []byte {
	return m.Bytes()
}
