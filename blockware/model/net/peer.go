package net

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/model/games"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

/**

A peer represents a single node in the network who will:
- own a collection of games
- interact with the ETH blockchain
- send and receive data from other peers

A peer is a singleton instance as only one will need to exist
per node.

*/

var (
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

	// config
	config PeerConfig

	// connections
	server  *TCPServer
	clients []*TCPClient

	// state
	peers map[PeerIT]*peerData

	// data
	installFolder string
	library       *games.Library

	// list of known peers addresses
	knownPeerAddresses []string
}

// Runtime configuration settings for the peer
type PeerConfig struct {

	// attempt to start downloads when starting the peer
	ContinueDownloads bool

	// read and attempt to connect to peers from a file
	LoadPeersFromFile bool
}

// Stores useful information about other peers
type peerData struct {

	// peer details for logging
	Hostname string
	Port     uint

	// socket interface to communicate with peer
	Peer PeerIT

	// what games they have installed
	Library map[[32]byte]bool
}

// Get the singleton instance of the current peer if it exists
func Peer() *peer {
	return singleton
}

// start a new instance of a peer and assign it to the singleton
func StartPeer(config PeerConfig, serverHostname string, serverPort uint, installFolder, toolkitFolder string) (*peer, error) {
	util.Logger.Info("Starting peer")
	once.Do(func() {
		peer, err := newPeer(config, serverHostname, serverPort, installFolder, toolkitFolder)
		if err != nil {
			util.Logger.Errorf("Error creating peer %s", err)
			return
		}

		singleton = peer
		if config.LoadPeersFromFile {
			go singleton.connectToKnownPeers()
		}

		if config.ContinueDownloads {
			singleton.listenToDownloadRequests()
		}
	})

	return singleton, nil
}

// create a new peer instance
func newPeer(config PeerConfig, serverHostname string, serverPort uint, installFolder, toolkitFolder string) (*peer, error) {
	gameLs, err := games.LoadGames(filepath.Join(toolkitFolder, "games"))
	if err != nil {
		return nil, err
	}
	util.Logger.Infof("Found %d games", len(gameLs))

	lib := games.NewLibrary()
	for _, g := range gameLs {
		err = lib.AddOwnedGame(g)
		if err != nil && !os.IsNotExist(err) {
			util.Logger.Error(err)
		}
	}

	lib.OutputGamesTable()

	var knownPeers []string = []string{}
	if config.LoadPeersFromFile {
		knownPeers, err = loadPeersFromFile()
		if err != nil {
			util.Logger.Warnf("Error fetching known peers from file %s", err)
		}
	}

	peer := &peer{
		server:             InitServer(serverHostname, serverPort),
		clients:            []*TCPClient{},
		peers:              make(map[PeerIT]*peerData),
		installFolder:      installFolder,
		library:            lib,
		knownPeerAddresses: knownPeers,
		config:             config,
	}

	go peer.server.Start(onMessage)
	return peer, nil
}

// form a connection to another peer
func (p *peer) ConnectToPeer(hostname string, portNo uint) error {
	client, err := InitTCPClient(hostname, portNo)
	if err != nil {
		return err
	}

	p.clients = append(p.clients, client)
	return nil
}

// run this function every time we connect to a new peer
func (p *peer) onConnection(hostname string, port uint, peer PeerIT) {
	p.peers[peer] = &peerData{
		Hostname: hostname,
		Port:     port,
		Peer:     peer,
		Library:  make(map[[32]byte]bool),
	}

	err := peer.SendString(generateLIBRARY())
	if err != nil {
		util.Logger.Warnf("Error sending message to peer %s: %s", peer.Info(), err)
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

// * known peers

// load peers from file containing peer info
func loadPeersFromFile() ([]string, error) {
	util.Logger.Info("Attempting to read peer list from file")
	file, err := os.Open(filepath.Join(viper.GetString("meta.directory"), "peers.txt"))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	out := []string{}

	var line string
	for {
		if line, err = reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			}

			util.Logger.Warn(err)
			continue
		}

		out = append(out, line[:len(line)-1])
	}

	util.Logger.Infof("Read %d peers from file", len(out))
	return out, nil
}

// attempts to form connections to the list of known peers
func (p *peer) connectToKnownPeers() {
	ps, err := loadPeersFromFile()
	if err != nil {
		util.Logger.Warnf("Error reading from peers file %s", err)
		return
	}

	count := 0
	for _, peer := range ps {
		peerInfo := strings.Split(peer, ":")

		// ! remove carriage return if it exists
		if peerInfo[1][len(peerInfo[1])-1] == '\r' {
			peerInfo[1] = peerInfo[1][:len(peerInfo[1])-1]
		}

		port, err := strconv.ParseUint(peerInfo[1], 10, 16)
		if err != nil {
			util.Logger.Warnf("Error parsing peer details for peer %s: %s", peer, err)
			continue
		}

		err = p.ConnectToPeer(peerInfo[0], uint(port))
		if err != nil {
			util.Logger.Warnf("Error conneting to peer %s: %s", err)
			continue
		}

		count++
	}

	util.Logger.Infof("Connecting to %d known peers", count)
}

// save current list of peers to file
func (p *peer) savePeersToFile() error {
	util.Logger.Info("Writing peers to file")
	file, err := os.Create(filepath.Join(viper.GetString("meta.directory"), "peers.txt"))
	if err != nil {
		return err
	}
	defer file.Close()

	for peer := range p.GetPeers() {
		file.WriteString(peer.Info() + "\n")
	}

	util.Logger.Info("Peers written to file")
	return nil
}

// * GETTERS

// get a list of peers
func (p *peer) GetPeers() map[PeerIT]*peerData {
	return p.peers
}

// get an existing peer
func (p *peer) GetPeer(peer PeerIT) *peerData {
	return p.peers[peer]
}

// Get the library of the current peer
func (p *peer) Library() *games.Library {
	return p.library
}

// * UTIL

// shutdown the peer
func (p *peer) Close() {
	util.Logger.Info("Closing down peer")
	p.savePeersToFile()
	p.library.Close()
	p.server.Close()
	for _, c := range p.clients {
		c.Close()
	}
}

// send a message to all peers
func (p *peer) Broadcast(message string) {
	for peer := range p.peers {
		peer.SendString(message)
	}
}

// * CONFIG settings

func (p *peer) SetContinueDownloads(state bool) {
	if p.config.ContinueDownloads == state {
		return
	}

	util.Logger.Infof("Setting continue downloads to %s", state)
	p.config.ContinueDownloads = state
	if state {
		// start downloads
		p.library.ContinueDownloads()
		singleton.listenToDownloadRequests()
	} else {
		// pause downloads
		p.library.StopDownloads()
	}
}
