package main

import (
	"github.com/agoussia/godes"
	"poasim/network"
	"fmt"
	"math/rand"
	"time"
)


var uniform = godes.NewUniformDistr(true)


/* kod koji će izvršavati pojedini node - validator */
func validatorCode(n *network.Node) {

	// Ako imam poruku u svom queue
	if n.HasMessage() {
		rcv_msg := n.PopMessage()
		fmt.Printf("[%f] %s\n", godes.GetSystemTime(), rcv_msg)

		peers := n.GetPeers() // daj mi listu peerova (nema discoverya, sve adrese su poznate za sada)
		id := peers[rand.Intn(len(peers))] // daj mi random peer-a
		peer := network.GetNodeById(id)
		msg := network.NewPingMessage(n.GetId(), peer.GetId()) // pošalji poruku preko network sloja
		network.SendMessage(msg)

	} else {
		// Spavaj i čekaj
		wait := uniform.Get(100, 500)
		godes.Advance(wait)
	}
}

func main() {

	rand.Seed(time.Now().Unix())
	peers := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, id := range peers {
		// super je što mogu predati funkciju definicije nodea (što radi) izvana.
		n := network.NewNode(id, peers, validatorCode)
		godes.AddRunner(n)
	}
	godes.Run()

	msg := network.NewPingMessage(1, 2)
	network.SendMessage(msg)

	godes.WaitUntilDone()
}
