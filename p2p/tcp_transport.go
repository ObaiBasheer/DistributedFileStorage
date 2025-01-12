package p2p

import (
	"net"
	"sync"
)

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	mu 		  sync.Mutex	
	peers		 map[string]*Peer
}