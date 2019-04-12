package network

import "fmt"

type Message struct {
	Type string
	Content string
	From int
	To int
}

func NewPingMessage(from int, to int) (p *Message) {
	p = &Message{}
	p.Type = "PING"
	p.Content = "ping"
	p.From = from
	p.To = to
	return
}

func (m *Message) String() string {
	return fmt.Sprintf("[%s] from %d to %d", m.Type, m.From, m.To)
}
