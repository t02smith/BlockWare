package peer

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/model/manager/games"
	"github.com/t02smith/part-iii-project/toolkit/model/net/tcp"
	model "github.com/t02smith/part-iii-project/toolkit/model/util"
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

	contributions *contributions

	// data
	installFolder string
	library       *games.Library

	// list of known peers addresses
	knownPeerAddresses []string
}

// Config Runtime configuration settings for the peer
type Config struct {

	// hostname or ip address of our server
	PublicHostname string

	// attempt to start downloads when starting the peer
	ContinueDownloads bool

	// read and attempt to connect to peers from a file
	LoadPeersFromFile bool

	//
	ServeAssets bool

	// whether to enforce address validation
	SkipValidation bool

	// track contributions sent by peers
	TrackContributions bool

	// upper limit on connections
	MaxConnections uint
}

// * functions

// Peer Get the singleton instance of the current peer if it exists
func Peer() *peer {
	return singleton
}

// reutrn the config settings for this peer
func (p *peer) Config() Config {
	return p.config
}

// StartPeer start a new instance of a peer and assign it to the singleton
func StartPeer(config Config, serverHostname string, serverPort uint, installFolder, toolkitFolder string) (*peer, error) {
	util.Logger.Info("Starting peer")
	once.Do(func() {
		if err := model.SetupToolkitEnvironment(); err != nil {
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
		knownPeers, err = LoadPeersFromFile(filepath.Join(viper.GetString("meta.directory"), "peers.txt"))
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
		contributions:      newContributions(),
	}

	go peer.server.Start(OnMessage, peer.onConnection, peer.OnConnectionClose)
	return peer, nil
}

// ConnectToPeer form a connection to another peer
func (p *peer) ConnectToPeer(hostname string, portNo uint) error {
	client, err := tcp.InitTCPClient(hostname, portNo, OnMessage, p.onConnection, p.OnConnectionClose)
	if err != nil {
		return err
	}

	p.clients = append(p.clients, client)
	return nil
}

// run this function every time we connect to a new peer
func (p *peer) onConnection(hostname string, port uint, peer tcp.TCPConnection) {
	numOfPeers := len(p.peers)

	if numOfPeers == int(p.config.MaxConnections) {
		util.Logger.Debug("max connections reached => rejecting connection")
		if err := peer.Close(); err != nil {
			util.Logger.Warn(err)
		}
		return
	}

	pd := &peerData{
		Hostname:             hostname,
		Port:                 port,
		Peer:                 peer,
		Library:              make(map[[32]byte]ownership),
		sentRequests:         make(map[games.DownloadRequest]time.Time, 100),
		TotalRequestsSent:    0,
		TotalRepliesReceived: 0,
		Server:               "",
	}

	if port != 0 {
		// we have connected to them
		pd.Server = fmt.Sprintf("%s:%d", hostname, port)

		// send our server information to them
		host, port := p.GetServerInfo()
		if host != "" {
			peer.SendString(generateSERVER(host, port))
		}
	}

	p.setPeerData(peer, pd)

	// ? start address verification handshake
	if !p.config.SkipValidation {
		pd.ValidatePeer()
	}
}

// run this function after closing a connection to an existing peer
func (p *peer) OnConnectionClose(peer tcp.TCPConnection) {
	p.peersMU.Lock()
	util.Logger.Infof("Closing connection to %s", peer.Info())
	if err := peer.Close(); err != nil {
		util.Logger.Warnf("Err closing connection %s", err)
	}

	var incompleteReqs []games.DownloadRequest

	pd, ok := p.peers[peer]
	if ok {
		pd.lock.Lock()
		for req := range pd.sentRequests {
			incompleteReqs = append(incompleteReqs, req)
		}
		pd.lock.Unlock()
	}

	p.peersMU.Unlock()
	p.DeletePeer(peer)
	util.Logger.Infof("Connection closed to %s", peer.Info())

	go func() {
		reqChan := p.library.DownloadManager.RequestDownload
		util.Logger.Debugf("queueing incomplete requests")
		for _, req := range incompleteReqs {
			reqChan <- req
		}
	}()
}

// GetServerInfo get information about the current peer
func (p *peer) GetServerInfo() (string, uint) {
	host := p.config.PublicHostname
	_, port := p.server.GetServerInfo()
	return host, port
}

// * known peers

// load peers from file containing peer info
func LoadPeersFromFile(path string) ([]string, error) {
	util.Logger.Info("Attempting to read peer list from file")
	file, err := os.Open(path)
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
	ps, err := LoadPeersFromFile(filepath.Join(viper.GetString("meta.directory"), "peers.txt"))
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
	file, err := os.OpenFile(filepath.Join(viper.GetString("meta.directory"), "peers.txt"), os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755)
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
	p.contributions.writeContributions()
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

func (p *peer) SetPublicHostname(hostname string) {
	p.config.PublicHostname = hostname
	viper.Set("net.hostname", hostname)

	_, port := p.GetServerInfo()
	p.Broadcast(generateSERVER(hostname, port))
}

func (p *peer) setPeerData(con tcp.TCPConnection, pd *peerData) {
	p.peersMU.Lock()
	defer p.peersMU.Unlock()

	p.peers[con] = pd
}

// update handlers
func (p *peer) SetOnMessage(onMessage func([]string, tcp.TCPConnection) error) {
	p.server.SetOnMessage(onMessage)
}
