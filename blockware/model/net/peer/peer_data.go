package peer

import (
	"net/http"

	"github.com/t02smith/part-iii-project/toolkit/model/net/tcp"
	"github.com/t02smith/part-iii-project/toolkit/model/persistence/ethereum"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

// Stores useful information about other peers
type peerData struct {

	// peer details for logging
	Hostname string
	Port     uint

	// socket interface to communicate with peer
	Peer tcp.TCPConnection

	// the peer's game library <hash, ownership verified?>
	Library map[[32]byte]bool

	// used to validate the peer's identity
	Validator *ethereum.AddressValidator
}

// find peers who have a given game in their library
func (p *peer) findPeersWhoHaveGame(gameHash [32]byte) []tcp.TCPConnection {
	peers := []tcp.TCPConnection{}
	for peerIT, peer := range p.peers {
		if _, ok := peer.Library[gameHash]; ok {
			peers = append(peers, peerIT)
		}
	}

	return peers
}

// serve game assets to be fetched from the frontend locally
func (p *peer) serveAssetsFolder() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	util.Logger.Info("Serving assets on port 3003")
	err := http.ListenAndServe(":3003", nil)
	if err != nil {
		util.Logger.Error(err)
	}
}
