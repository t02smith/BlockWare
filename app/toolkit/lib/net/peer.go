package net

import (
	"fmt"

	"github.com/t02smith/part-iii-project/toolkit/cmd/view"
	"github.com/t02smith/part-iii-project/toolkit/lib/games"
)

type PeerIT interface {
	Send(command []byte) error
	SendString(command string) error
}

type Peer struct {

	// connections
	server  *TCPServer
	clients []*TCPClient

	// data
	installFolder string
	games         []*games.Game
}

type PeerData struct {
	Hostname string
	Port     uint
	Peer     PeerIT
}

func StartPeer(serverHostname string, serverPort uint, installFolder, gameDataLocation string) (*Peer, error) {
	gameLs, err := games.LoadGames(gameDataLocation)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Found %d games\n", len(gameLs))
	view.OutputGamesTable(gameLs)

	p := &Peer{
		server:        InitServer(serverHostname, serverPort),
		clients:       []*TCPClient{},
		installFolder: installFolder,
		games:         gameLs,
	}

	go p.server.Start(onMessage)
	return p, nil
}

// commands

func (p *Peer) ConnectToPeer(hostname string, portNo uint) error {

	client, err := InitTCPClient(hostname, portNo)
	if err != nil {
		return err
	}

	p.clients = append(p.clients, client)
	return nil
}

func (p *Peer) GetPeers() []PeerData {

	peers := []PeerData{}

	for _, p := range p.clients {
		peers = append(peers, PeerData{
			Hostname: p.hostname,
			Port:     p.port,
			Peer:     p,
		})
	}

	for _, p := range p.server.clients {
		peers = append(peers, PeerData{
			Hostname: p.con.LocalAddr().String(),
			Port:     0000,
			Peer:     p,
		})
	}

	return peers
}
