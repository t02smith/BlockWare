package peer

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/model"
	"github.com/t02smith/part-iii-project/toolkit/model/manager/games"
	"github.com/t02smith/part-iii-project/toolkit/model/net/tcp"
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

// * types

// A peer represents a node in a distributed network
type peer struct {

	// config
	config Config

	// connections
	server  *tcp.TCPServer
	clients []*tcp.TCPClient

	// state
	peers   map[tcp.TCPConnection]*peerData
	peersMU sync.Mutex

	// data
	installFolder string
	library       *games.Library

	// list of known peers addresses
	knownPeerAddresses []string
}

// Config Runtime configuration settings for the peer
type Config struct {

	// attempt to start downloads when starting the peer
	ContinueDownloads bool

	// read and attempt to connect to peers from a file
	LoadPeersFromFile bool

	//
	ServeAssets bool

	// whether to enforce address validation
	SkipValidation bool
}

// * functions

// Peer Get the singleton instance of the current peer if it exists
func Peer() *peer {
	return singleton
}

// StartPeer start a new instance of a peer and assign it to the singleton
func StartPeer(config Config, serverHostname string, serverPort uint, installFolder, toolkitFolder string) (*peer, error) {
	util.Logger.Info("Starting peer")
	once.Do(func() {
		err := model.SetupToolkitEnvironment()
		if err != nil {
			util.Logger.Errorf("Error setting up toolkit dir: %s", err)
			return
		}

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

		if config.ServeAssets {
			go peer.serveAssetsFolder()

		}
	})

	return singleton, nil
}

// create a new peer instance
func newPeer(config Config, serverHostname string, serverPort uint, installFolder, toolkitFolder string) (*peer, error) {
	gameLs, err := games.LoadGames(filepath.Join(toolkitFolder, "games"))
	if err != nil {
		return nil, err
	}
	util.Logger.Infof("Found %d games", len(gameLs))

	lib := games.NewLibrary()
	for _, g := range gameLs {
		lib.AddOrUpdateOwnedGame(g)
	}

	var knownPeers []string
	if config.LoadPeersFromFile {
		knownPeers, err = loadPeersFromFile()
		if err != nil {
			util.Logger.Warnf("Error fetching known peers from file %s", err)
		}
	}

	peer := &peer{
		server:             tcp.InitServer(serverHostname, serverPort),
		clients:            []*tcp.TCPClient{},
		peers:              make(map[tcp.TCPConnection]*peerData),
		installFolder:      installFolder,
		library:            lib,
		knownPeerAddresses: knownPeers,
		config:             config,
	}

	go peer.server.Start(onMessage, peer.onConnection, peer.OnConnectionClose)
	return peer, nil
}

// ConnectToPeer form a connection to another peer
func (p *peer) ConnectToPeer(hostname string, portNo uint) error {
	client, err := tcp.InitTCPClient(hostname, portNo, onMessage, p.onConnection, p.OnConnectionClose)
	if err != nil {
		return err
	}

	p.clients = append(p.clients, client)
	return nil
}

// run this function every time we connect to a new peer
func (p *peer) onConnection(hostname string, port uint, peer tcp.TCPConnection) {
	pd := &peerData{
		Hostname:     hostname,
		Port:         port,
		Peer:         peer,
		Library:      make(map[[32]byte]ownership),
		sentRequests: make(map[games.DownloadRequest]model.Void, maxRequestsPerPeer),
	}
	p.setPeerData(peer, pd)

	// ? start address verification handshake
	if !p.config.SkipValidation {
		pd.validatePeer()
	}
}

// run this function after closing a connection to an existing peer
func (p *peer) OnConnectionClose(peer tcp.TCPConnection) {
	util.Logger.Infof("Closing connection to %s", peer.Info())
	peer.Close()
	p.DeletePeer(peer)
}

// GetServerInfo get information about the current peer
func (p *peer) GetServerInfo() (string, uint) {
	return "", 0
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

// GetPeers get a list of peers
func (p *peer) GetPeers() map[tcp.TCPConnection]*peerData {
	return p.peers
}

// GetPeer get an existing peer
func (p *peer) GetPeer(peer tcp.TCPConnection) *peerData {
	p.peersMU.Lock()
	defer p.peersMU.Unlock()

	return p.peers[peer]
}

func (p *peer) DeletePeer(peer tcp.TCPConnection) {
	p.peersMU.Lock()
	defer p.peersMU.Unlock()

	delete(p.peers, peer)
}

// Library Get the library of the current peer
func (p *peer) Library() *games.Library {
	return p.library
}

// * UTIL

// Close shutdown the peer
func (p *peer) Close() {
	util.Logger.Info("Closing down peer")
	p.savePeersToFile()
	p.library.Close()
	p.server.Close()
	for _, c := range p.clients {
		c.Close()
	}
}

// Broadcast send a message to all peers
func (p *peer) Broadcast(message string) {
	p.peersMU.Lock()
	defer p.peersMU.Unlock()

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

func (p *peer) setPeerData(con tcp.TCPConnection, pd *peerData) {
	p.peersMU.Lock()
	defer p.peersMU.Unlock()

	p.peers[con] = pd
}
