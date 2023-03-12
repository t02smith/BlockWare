package peer

import (
	"net/http"
	"sort"

	"github.com/t02smith/part-iii-project/toolkit/model"
	"github.com/t02smith/part-iii-project/toolkit/model/manager/games"
	"github.com/t02smith/part-iii-project/toolkit/model/net/tcp"
	"github.com/t02smith/part-iii-project/toolkit/model/persistence/ethereum"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

type ownership int8

const (
	owned ownership = iota
	notOwned
	unknown
)

// Stores useful information about other peers
type peerData struct {

	// peer details for logging
	Hostname string
	Port     uint

	// socket interface to communicate with peer
	Peer tcp.TCPConnection

	// the peer's game library <hash, ownership verified?>
	Library map[[32]byte]ownership

	// used to validate the peer's identity
	Validator *ethereum.AddressValidator

	//
	sentRequests map[games.DownloadRequest]model.Void
}

// validate the identity of a given peer
func (pd *peerData) validatePeer() {
	if pd.Validator != nil && pd.Validator.Valid() {
		return
	}

	pd.Validator = ethereum.GenerateAddressValidation()
	pd.Peer.SendString(generateVALIDATE_REQ(pd.Validator.Message()))
}

func (pd *peerData) checkOwnership(gameHash [32]byte) (bool, error) {
	// ? has the user's address been validated
	if pd.Validator == nil || !pd.Validator.Valid() {
		pd.validatePeer()
		return false, nil
	}

	checked, ok := pd.Library[gameHash]
	if !ok {
		return false, nil
	}

	switch checked {
	case owned:
		return true, nil
	case notOwned:
		return false, nil
	case unknown:
		return true, nil
		// addr := crypto.PubkeyToAddress(*pd.Validator.PublicKey)
		// verified, err := library.HasPurchased(gameHash, addr)
		// if err != nil {
		// 	return false, err
		// }

		// if verified {
		// 	pd.Library[gameHash] = owned
		// } else {
		// 	pd.Library[gameHash] = notOwned
		// }

		// return verified, nil
	}

	return false, nil
}

// find peers who have a given game in their library
func (p *peer) findPeersWhoHaveGame(gameHash [32]byte) []tcp.TCPConnection {
	var peers []tcp.TCPConnection

	p.peersMU.Lock()
	defer p.peersMU.Unlock()
	for peerIT, peer := range p.peers {
		if _, ok := peer.Library[gameHash]; ok {
			peers = append(peers, peerIT)
		}
	}

	// prioritise connections who have been sent the least blocks
	sort.Slice(peers, func(i, j int) bool {
		return len(p.peers[peers[i]].sentRequests) < len(p.peers[peers[j]].sentRequests)
	})

	return peers
}

// serve game assets to be fetched from the frontend locally
func (p *peer) serveAssetsFolder() {
	fs := http.FileServer(http.Dir("."))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		fs.ServeHTTP(w, r)
	})

	util.Logger.Info("Serving assets on port 3003")
	err := http.ListenAndServe(":3003", nil)
	if err != nil {
		util.Logger.Error(err)
	}
}
