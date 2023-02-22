package net

// find peers who have a given game in their library
func (p *peer) findPeersWhoHaveGame(gameHash [32]byte) []PeerIT {
	peers := []PeerIT{}
	for peerIT, peer := range p.peers {
		if _, ok := peer.Library[gameHash]; ok {
			peers = append(peers, peerIT)
		}
	}

	return peers
}
