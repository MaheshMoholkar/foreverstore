package main

import (
	"fmt"
	"log"

	"github.com/MaheshMoholkar/foreverstore/p2p"
)

func OnPeer(p p2p.Peer) error {
	fmt.Println("onpeer func executing")
	return nil
}

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        OnPeer,
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
