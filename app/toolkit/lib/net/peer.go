package net

type Peer struct {

	// connections
	server  *TCPServer
	clients []*TCPClient

	// data
	softwareFolder string
}
