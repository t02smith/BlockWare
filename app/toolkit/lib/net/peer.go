package net

import "fmt"

type Peer struct {

	// connections
	server  *TCPServer
	clients []*TCPClient

	// data
	installFolder string
}

func StartPeer(serverHostname string, serverPort uint, installFolder string) {
	p := &Peer{
		server:        InitServer(serverHostname, serverPort),
		clients:       []*TCPClient{},
		installFolder: installFolder,
	}

	go p.server.Start(onMessage)

}

//

func onMessage(cmd []byte, client *TCPServerClient) {
	switch cmd[0] {

	// Library
	case 0x01:
		fmt.Println("Library command called")
	}
}
