package net

import (
	"sync"

	"github.com/t02smith/part-iii-project/toolkit/model/games"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

var (
	// peer is a singleton type so should only be instantiated once
	singleton *peer
	once      sync.Once
)

// represents a peer that we are connected to
// this is an abstraction as a peer may connect to our server socket or
// we might connect to theirs
type PeerIT interface {
	Send(command []byte) error
	SendString(command string) error
	SendStringf(command string, args ...any) error
	Info() string
}

// A peer represents a node in a distributed network
type peer struct {

	// connections
	server  *TCPServer
	clients []*TCPClient

	// state
	peers map[PeerIT]*PeerData

	// data
	installFolder string
	library       *games.Library
}

// Stores useful information about other peers
type PeerData struct {

	// peer details for logging
	Hostname string
	Port     uint

	// socket interface to communicate with peer
	Peer PeerIT

	// what games they have installed
	Library []*games.Game
}

// Get the singleton instance of the current peer if it exists
func GetPeerInstance() *peer {
	return singleton
}

// start a new instance of a peer
func StartPeer(serverHostname string, serverPort uint, installFolder, gameDataLocation string) (*peer, error) {
	util.Logger.Info("Starting peer")
	once.Do(func() {
		gameLs, err := games.LoadGames(gameDataLocation)
		if err != nil {
			return
		}

		util.Logger.Infof("Found %d games\n", len(gameLs))

		lib := games.NewLibrary()
		for _, g := range gameLs {
			lib.AddGame(g)
		}

		lib.OutputGamesTable()

		singleton = &peer{
			server:        InitServer(serverHostname, serverPort),
			clients:       []*TCPClient{},
			peers:         make(map[PeerIT]*PeerData),
			installFolder: installFolder,
			library:       lib,
		}

		go singleton.server.Start(onMessage)
	})

	return singleton, nil
}

// run this function every time we connect to a new peer
func (p *peer) onConnection(hostname string, port uint, peer PeerIT) {
	p.peers[peer] = &PeerData{
		Hostname: hostname,
		Port:     port,
		Peer:     peer,
		Library:  []*games.Game{},
	}
}

// run this function after closing a connection to an existing peer
func (p *peer) onClose(peer PeerIT) {
	util.Logger.Infof("Closing connection to %s", peer.Info())
	delete(p.peers, peer)
}

// get information about the current peer
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

// send a message to all peers
func (p *peer) Broadcast(message string) {
	for peer := range p.peers {
		peer.SendString(message)
	}
}

func (p *peer) GetLibrary() *games.Library {
	return p.library
}

// shutdown the peer
func (p *peer) Close() {
	p.server.Close()
	for _, c := range p.clients {
		c.Close()
	}
}
