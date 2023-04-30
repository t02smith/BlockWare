package peer

import (
	"math/rand"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/t02smith/part-iii-project/toolkit/model/manager/games"
	"github.com/t02smith/part-iii-project/toolkit/model/net/tcp"
	"github.com/t02smith/part-iii-project/toolkit/model/persistence/ethereum"
	"github.com/t02smith/part-iii-project/toolkit/model/persistence/ethereum/library"

	"github.com/t02smith/part-iii-project/toolkit/util"
)

type ownership int8

const (
	owned ownership = iota
	notOwned
	unknown
)

// rank a peer higher at this rate (%) to unchoke downloads
const unchokeRate int8 = 25

// Stores useful information about other peers
type peerData struct {

	// peer details for logging
	Hostname string
	Port     uint
	Server   string

	// socket interface to communicate with peer
	Peer tcp.TCPConnection

	// the peer's game library <hash, ownership verified?>
	Library map[[32]byte]ownership

	// used to validate the peer's identity
	Validator *ethereum.AddressValidator

	//
	sentRequests         map[games.DownloadRequest]time.Time
	TotalRequestsSent    int64
	TotalRepliesReceived int64

	lock sync.Mutex
}

// validate the identity of a given peer
func (pd *peerData) ValidatePeer() {
	if pd.Validator != nil && pd.Validator.Valid() {
		return
	}

	pd.Validator = ethereum.GenerateAddressValidation()
	pd.Peer.SendString(generateVALIDATE_REQ(pd.Validator.Message()))
}

func (pd *peerData) checkOwnership(gameHash [32]byte) (bool, error) {
	// ? has the user's address been validated
	if pd.Validator == nil || !pd.Validator.Valid() {
		pd.ValidatePeer()
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
		util.Logger.Debugf("Verifying ownership of game %x for user %s", gameHash, pd.Peer.Info())
		addr := crypto.PubkeyToAddress(*pd.Validator.PublicKey)
		verified, err := library.HasPurchased(gameHash, addr)
		if err != nil {
			return false, err
		}

		if verified {
			util.Logger.Debugf("User %s owns game %x", pd.Peer.Info(), gameHash)
			pd.Library[gameHash] = owned
		} else {
			util.Logger.Debugf("User %s does not own game %x", pd.Peer.Info(), gameHash)
			pd.Library[gameHash] = notOwned
		}

		return verified, nil
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

	r := rand.New(rand.NewSource(time.Now().Unix()))
	sort.Slice(peers, func(i, j int) bool {
		p1 := p.peers[peers[i]]
		p2 := p.peers[peers[j]]

		// always send to unkown peers first to help discovery
		if p1.TotalRequestsSent == 0 && p2.TotalRequestsSent > 0 ||
			p1.TotalRequestsSent == 0 && p2.TotalRequestsSent == 0 {
			return true
		} else if p1.TotalRequestsSent > 0 && p2.TotalRequestsSent == 0 {
			return false
		}

		// otherwise send to peer with highest reputation
		rep1 := float32(p1.TotalRepliesReceived) / float32(p1.TotalRequestsSent)
		rep2 := float32(p2.TotalRepliesReceived) / float32(p2.TotalRequestsSent)
		unchoke := r.Intn(100) >= int(unchokeRate)

		return rep1 <= rep2 && !unchoke
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

// store details about a request sent to a user
func (pd *peerData) PushRequest(req games.DownloadRequest) {
	pd.lock.Lock()
	defer pd.lock.Unlock()

	if _, ok := pd.sentRequests[req]; ok {
		util.Logger.Debug("req already sent to peer => ignoring")
		return
	}

	pd.sentRequests[req] = time.Now()
	pd.TotalRequestsSent++
}

// confirm that a request was replied to
func (pd *peerData) ConfirmRequest(req games.DownloadRequest) {
	pd.lock.Lock()
	defer pd.lock.Unlock()

	if _, ok := pd.sentRequests[req]; !ok {
		return
	}

	delete(pd.sentRequests, req)
	pd.TotalRepliesReceived++
}

// confirm a request was not replied to
func (pd *peerData) FailedRequest(req games.DownloadRequest) {
	pd.lock.Lock()
	defer pd.lock.Unlock()

	if _, ok := pd.sentRequests[req]; !ok {
		return
	}

	delete(pd.sentRequests, req)
}

func (pd *peerData) Lock() {
	pd.lock.Lock()
}

func (pd *peerData) Unlock() {
	pd.lock.Unlock()
}
