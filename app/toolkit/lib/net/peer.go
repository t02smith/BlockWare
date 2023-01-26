package net

import (
	"log"
	"sync"

	"github.com/t02smith/part-iii-project/toolkit/cmd/view"
	"github.com/t02smith/part-iii-project/toolkit/lib/games"
)

var (
	singleton *peer
	once      sync.Once
)

type PeerIT interface {
	Send(command []byte) error
	SendString(command string) error
}

type peer struct {

	// connections
	server  *TCPServer
	clients []*TCPClient

	// state
	// peers []*PeerData

	peers map[PeerIT]*PeerData

	// data
	installFolder string
	games         []*games.Game
}

type PeerData struct {
	Hostname string
	Port     uint
	Peer     PeerIT
	Library  []*games.Game
}

func GetPeerInstance() *peer {
	return singleton
}

func StartPeer(serverHostname string, serverPort uint, installFolder, gameDataLocation string) (*peer, error) {
	log.Println("Starting peer")
	once.Do(func() {
		gameLs, err := games.LoadGames(gameDataLocation)
		if err != nil {
			return
		}

		log.Printf("Found %d games\n", len(gameLs))
		view.OutputGamesTable(gameLs)

		singleton = &peer{
			server:        InitServer(serverHostname, serverPort),
			clients:       []*TCPClient{},
			peers:         make(map[PeerIT]*PeerData),
			installFolder: installFolder,
			games:         gameLs,
		}

		go singleton.server.Start(onMessage)
	})

	return singleton, nil
}

func (p *peer) onConnection(hostname string, port uint, peer PeerIT) {
	p.peers[peer] = &PeerData{
		Hostname: hostname,
		Port:     port,
		Peer:     peer,
		Library:  []*games.Game{},
	}
}

func (p *peer) onClose(peer PeerIT) {
	delete(p.peers, peer)
}

func (p *peer) GetServerInfo() (string, uint) {
	return p.server.hostname, p.server.port
}

// commands

// form a connection to another peer
func (p *peer) ConnectToPeer(hostname string, portNo uint) error {

	client, err := InitTCPClient(hostname, portNo)
	if err != nil {
		return err
	}

	p.clients = append(p.clients, client)
	return nil
}

// get a list of peers
func (p *peer) GetPeers() map[PeerIT]*PeerData {
	return p.peers
}

// get an existing peer
func (p *peer) GetPeer(peer PeerIT) *PeerData {
	return p.peers[peer]
}
