package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represent the remote node over TCP established connection.
type TCPPeer struct {
	//conn the underline connection of  the peer
	conn net.Conn

	//if we dial and retrieve a conn => outbound == true
	//if we accept and retrevie a conn => outbound == true

	outbound bool
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	shakeHand     HandShakeFunc
	decoder       Decoder
	mu            sync.Mutex
	peers         map[string]*Peer
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		shakeHand:     NOPHandshakeFunc,
		listenAddress: listenAddress,
		//peers: make(map[string]*Peer),
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)

	if err != nil {
		return err
	}

	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
			continue
		}
		fmt.Println("New connection accepted:", conn.RemoteAddr())
		go t.handleConn(conn)
	}
}

type Temp struct{}

func (t *TCPTransport) handleConn(cnn net.Conn) {
	peer := TCPPeer{conn: cnn, outbound: true}
	if err := t.shakeHand(peer); err != nil {
		cnn.Close()
		fmt.Printf("TCP handshake Error : %s\n", err)
	}

	//buf := new(bytes.Buffer)

	msg := &Temp{}
	for {
		if err := t.decoder.Decode(cnn, msg); err != nil {
			fmt.Printf("TCP Error: %s\n", err)
			continue
		}
	}

	//fmt.Printf("handle a new connection %+v\n", peer)
}
