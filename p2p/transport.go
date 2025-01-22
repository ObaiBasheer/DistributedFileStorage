package p2p

// Peer is an interface that represents a peer in the network (a remote note).
type Peer interface{}

// Transport is an interface that handles the communication between peers. This can be a TCP connection, a UDP connection, a WebRTC connection, etc.
type Transport interface {
	ListenAndAccept() error
}
