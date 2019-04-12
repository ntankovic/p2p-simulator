package network

import (
	"github.com/agoussia/godes"
	"fmt"
)

type INode interface {
	HasMessage()
	AddMessage()
	Run()
}

type RunFunction func(*Node)

type Node struct {
	*godes.Runner
	id int
	queue *godes.FIFOQueue
	peers []int
	runFunction RunFunction
	TotalMessages int
}

var TotalMessages = 0
var nodes = make(map[int]*Node)
var uniform = godes.NewUniformDistr(true)

func GetNodeById(id int) (*Node) {
	return nodes[id]
}

func NewNode(id int, peers []int, fn RunFunction) (n* Node) {
	n = new(Node)
	n.Runner = &godes.Runner{}
	n.id = id
	n.queue = godes.NewFIFOQueue("messages")
	n.peers = peers
	n.runFunction = fn
	nodes[id] = n
	fmt.Printf("Added nodes %d with peers %+v\n", n.id, n.peers)
	return
}

func (n *Node) GetId() int {
	return n.id
}
func (n *Node) GetPeers() ([]int) {
	return n.peers
}

func (n *Node) HasMessage() bool {
	return n.queue.Len() > 0
}
func (n *Node) PopMessage() (m *Message) {
	TotalMessages++
	n.TotalMessages++
	return n.queue.Get().(*Message)
}
func (n *Node) AddMessage(m *Message) {
	n.queue.Place(m)
}

func (n *Node) Run() {
	fmt.Println("Running...")

	for {
		n.runFunction(n)
	}
}

