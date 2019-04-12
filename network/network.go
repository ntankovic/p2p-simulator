package network

import (
	"github.com/agoussia/godes"
)

var networkLatency *godes.ExpDistr = godes.NewExpDistr(true)

func SendMessage(msg *Message) {
	godes.Advance(networkLatency.Get(100))
	destination := GetNodeById(msg.To)
	destination.AddMessage(msg)
}
