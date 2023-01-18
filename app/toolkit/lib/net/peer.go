package net

import (
	"fmt"

	"github.com/t02smith/part-iii-project/toolkit/lib/games"
)

type Peer struct {

	// connections
	server  *TCPServer
	clients []*TCPClient

	// data
	installFolder string
	games         []*games.Game
}

func StartPeer(serverHostname string, serverPort uint, installFolder, gameDataLocation string) error {
	gameLs, err := games.LoadGames(gameDataLocation)
	if err != nil {
		return err
	}

	p := &Peer{
		server:        InitServer(serverHostname, serverPort),
		clients:       []*TCPClient{},
		installFolder: installFolder,
		games:         gameLs,
	}

	go p.server.Start(onMessage)
	return nil
}

//

func onMessage(cmd []byte, client *TCPServerClient) {
	switch cmd[0] {

	// Library
	case 0x01:
		fmt.Println("Library command called")
	}
}
