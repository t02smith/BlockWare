package peer

import (
	"bytes"
	"compress/flate"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/t02smith/part-iii-project/toolkit/model/manager/games"
	"github.com/t02smith/part-iii-project/toolkit/model/net/tcp"
	"github.com/t02smith/part-iii-project/toolkit/model/persistence/ethereum"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*

This file defines specific interactions/commands. For each command it:
1. has a generator function for creating a structured message given the
   required set of parameters. This reduces code duplication.
2. has a handler to deal with an incoming message from another peer,

Current handlers are:
LIBRARY => request a users game library
GAMES => receive a users game library
BLOCK => request a block of data from a user
SEND_BLOCK => receive a block of data from a user
ERROR => an error message to send to a peer
				 TODO properly deal with incoming errors
VALIDATE_REQ => request a signed message from a peer
VALIDATE_RES => respond to a validation request

*/

func getHandlerByCommand(cmd string) func([]string, tcp.TCPConnection) error {
	switch cmd {
	case "LIBRARY":
		return handleLIBRARY
	case "GAMES":
		return handleGAMES
	case "BLOCK":
		return handleBLOCK
	case "SEND_BLOCK":
		return handleSEND_BLOCK
	case "ERROR":
		return handleERROR
	case "VALIDATE_REQ":
		return handleVALIDATE_REQ
	case "VALIDATE_RES":
		return handleVALIDATE_RES
	case "REQ_RECEIPT":
		return handleREQ_RECEIPT
	case "RECEIPT":
		return handleRECEIPT
	case "REQ_PEERS":
		return handleREQ_PEERS
	case "PEERS":
		return handlePEERS
	}

	return func(s []string, t tcp.TCPConnection) error {
		return fmt.Errorf("unknown cmd '%s'", cmd)
	}
}

// LIBRARY

func generateLIBRARY() string {
	return "LIBRARY\n"
}

func handleLIBRARY(cmd []string, client tcp.TCPConnection) error {
	util.Logger.Debugf("Library command called")
	return client.SendString(generateGAMES(Peer().library.GetOwnedGames()...))
}

// GAMES <game hash>, <game hash>, ...

func generateGAMES(games ...*games.Game) string {
	var buf bytes.Buffer
	buf.WriteString("GAMES;")

	for i, g := range games {
		_, err := buf.WriteString(fmt.Sprintf("%x", g.RootHash))
		if err != nil {
			util.Logger.Warnf("Errors writing root hash for game %x => skipping from message", g.RootHash)
			continue
		}

		if i != len(games)-1 {
			buf.WriteString(";")
		}
	}

	buf.WriteString("\n")
	return buf.String()
}

func handleGAMES(cmd []string, client tcp.TCPConnection) error {
	util.Logger.Debugf("Games command called")
	changes := false

	pData := Peer().GetPeer(client)
	if pData == nil {
		return errors.New("unknown peer")
	}

	for i := 1; i < len(cmd); i++ {
		if len(cmd[i]) != 64 {
			continue
		}

		var hash [32]byte
		receivedHash, err := hex.DecodeString(cmd[i])
		if err != nil {
			return err
		}
		copy(hash[:], receivedHash)

		if _, ok := pData.Library[hash]; !ok {
			changes = true
			pData.Library[hash] = unknown
		}
	}

	// if new games are detected we can start trying to process
	// deferred requests
	if changes {
		LoadDeferredRequests()
	}

	return nil
}

// BLOCK <game hash> <block hash>

func generateBLOCK(gameHash, blockHash [32]byte) string {
	return fmt.Sprintf("BLOCK;%x;%x\n", gameHash, blockHash)
}

func handleBLOCK(cmd []string, client tcp.TCPConnection) error {
	util.Logger.Debugf("Block command called for block %s", cmd[2])

	gh, err := stringTo32ByteArr(cmd[1])
	if err != nil {
		return fmt.Errorf("error reading game hash on BLOCK cmd: %s", err)
	}

	sh, err := stringTo32ByteArr(cmd[2])
	if err != nil {
		return fmt.Errorf("error reading shard hash on BLOCK cmd: %s", err)
	}

	if !Peer().config.SkipValidation {
		pd := Peer().peers[client]
		ownsGame, err := pd.checkOwnership(gh)
		if err != nil {
			return err
		}

		if !ownsGame {
			return fmt.Errorf("user does not own game %x", gh)
		}
	}

	found, data, err := Peer().library.FindAndRetrieveBlock(gh, sh)
	if err != nil {
		return fmt.Errorf("error finding block %s", err)
	}

	if !found {
		client.SendStringf(generateERROR("Block %x not found", sh))
		return fmt.Errorf("block %x not found", sh)
	}

	client.SendString(generateSEND_BLOCK(gh, sh, data))
	return nil
}

// SEND_BLOCK <game hash> <shard hash> <shard data>

func generateSEND_BLOCK(gameHash, shardHash [32]byte, data []byte) string {
	// compress data
	var b bytes.Buffer
	w, _ := flate.NewWriter(&b, 6)
	w.Write(data)
	w.Close()

	return fmt.Sprintf("SEND_BLOCK;%x;%x;%x\n", gameHash, shardHash, b.Bytes())
}

func handleSEND_BLOCK(cmd []string, client tcp.TCPConnection) error {
	util.Logger.Debugf("SEND_BLOCK => Block received")
	// * parse input
	gh, err := stringTo32ByteArr(cmd[1])
	if err != nil {
		return fmt.Errorf("error reading game hash on BLOCK cmd: %s", err)
	}

	sh, err := stringTo32ByteArr(cmd[2])
	if err != nil {
		return fmt.Errorf("error reading shard hash on BLOCK cmd: %s", err)
	}

	// * fetch game and find location
	game := Peer().library.GetOwnedGame(gh)
	if game == nil {
		return fmt.Errorf("user does not own game with hash %x", gh)
	}

	download := game.GetDownload()
	if download == nil {
		return fmt.Errorf("download not found for %x", game.RootHash)
	}

	gameTree, err := game.GetData()
	if err != nil {
		return fmt.Errorf("error loading game data %s", err)
	}

	compressedData, err := hex.DecodeString(cmd[3])
	if err != nil {
		return err
	}

	// decompress
	var b bytes.Buffer
	w := flate.NewReader(bytes.NewReader(compressedData))
	b.ReadFrom(w)
	w.Close()

	data := b.Bytes()

	// ? data correct length
	if uint(len(data)) != gameTree.ShardSize {
		return fmt.Errorf("incorrect data length. expected %d, got %d", gameTree.ShardSize, len(data))
	}

	// ? find file
	locations := gameTree.FindShard(sh)
	if len(locations) == 0 {
		return fmt.Errorf("shard %x not found", sh)
	}

	for _, l := range locations {
		download.InserterPool() <- games.InsertShardRequest{
			FileHash:  l.File.RootHash,
			BlockHash: sh,
			Data:      data,
		}
	}

	pd := Peer().GetPeer(client)
	pd.Lock()
	delete(pd.sentRequests, games.DownloadRequest{
		GameHash:  gh,
		BlockHash: sh,
	})
	pd.Unlock()

	if pd.Validator != nil {
		Peer().contributions.addContribution(crypto.PubkeyToAddress(*pd.Validator.PublicKey), gh, sh)
	}

	game.OutputToFile()

	util.Logger.Debugf("Successfully inserted shard %x", sh)
	return nil
}

// ERROR <msg>

func generateERROR(msg string, args ...any) string {
	return fmt.Sprintf("ERROR;%s\n", fmt.Sprintf(msg, args...))
}

func handleERROR(cmd []string, tcp tcp.TCPConnection) error {
	util.Logger.Errorf("Error received %s", cmd[1])
	return nil
}

// VALIDATE_REQ <message>

func generateVALIDATE_REQ(message []byte) string {
	return fmt.Sprintf("VALIDATE_REQ;%x\n", message)
}

func handleVALIDATE_REQ(cmd []string, client tcp.TCPConnection) error {
	message, err := hex.DecodeString(cmd[1])
	if err != nil {
		return err
	}

	sig, err := ethereum.ProduceAddressValidation(message)
	if err != nil {
		return err
	}

	client.SendString(generateVALIDATE_RES(sig))
	return nil
}

// VALIDATE_RES <signed message>

func generateVALIDATE_RES(sig []byte) string {
	return fmt.Sprintf("VALIDATE_RES;%x\n", sig)
}

func handleVALIDATE_RES(cmd []string, client tcp.TCPConnection) error {
	sig, err := hex.DecodeString(cmd[1])
	if err != nil {
		return err
	}

	validator := Peer().GetPeer(client).Validator
	valid, err := ethereum.CheckAddressValidation(validator, sig)
	if err != nil {
		return err
	}

	if !valid {
		client.SendStringf(generateERROR("invalid signature sent"))
	} else {
		// request their libraries and a list of peers on connection
		client.SendString(generateLIBRARY())
		client.SendString(generateREQ_PEERS())
	}

	return nil
}

// REQ_RECEIPT

func generateREQ_RECEIPT(game [32]byte) string {
	return fmt.Sprintf("REQ_RECEIPT;%x\n", game)
}

func handleREQ_RECEIPT(cmd []string, client tcp.TCPConnection) error {
	p := Peer()

	gameHash, err := stringTo32ByteArr(cmd[1])
	if err != nil {
		return err
	}

	// 1 flush all contributions to file
	p.contributions.writeContributions()

	// 2 read contributions from user
	pd := p.GetPeer(client)
	if pd.Validator == nil {
		util.Logger.Warnf("peer not validated => cannot generate receipt")
		return nil
	}
	addr := crypto.PubkeyToAddress(*pd.Validator.PublicKey)

	blocks, err := GetContributionsFromPeer(addr, gameHash)
	if err != nil {
		return err
	}

	client.SendString(generateRECEIPT(gameHash, blocks))
	return nil
}

//

func generateRECEIPT(game [32]byte, blocks [][32]byte) string {
	data := make([]byte, 32*len(blocks))

	for i, b := range blocks {
		for j := 0; j < 32; j++ {
			data[(i*32)+j] = b[j]
		}
	}

	sig, _ := ethereum.SignMessage(data)
	return fmt.Sprintf("RECEIPT;%x;%x;%x\n", game, sig, data)
}

func handleRECEIPT(cmd []string, client tcp.TCPConnection) error {
	game, err := stringTo32ByteArr(cmd[1])
	if err != nil {
		return err
	}
	_ = game

	fmt.Printf("'%s'", cmd[3])
	data, err := hex.DecodeString(cmd[3])
	if err != nil {
		return err
	}

	// check signature of data matches expected user
	pd := Peer().peers[client]
	if pd.Validator == nil || !pd.Validator.Valid() {
		return fmt.Errorf("user not verified. ignoring receipt")
	}

	sig, err := hex.DecodeString(cmd[2])
	if err != nil {
		return err
	}

	pubKeyBytes := crypto.FromECDSAPub(pd.Validator.PublicKey)
	hash := crypto.Keccak256Hash(data)

	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), sig)
	if err != nil {
		return err
	}

	if !bytes.Equal(pubKeyBytes, sigPublicKey) {
		return fmt.Errorf("public keys do not match => throwing receipt")
	}

	util.Logger.Infof("Received receipt for %d blocks", len(data)/32)
	// TODO do something ??

	return nil
}

// REQ_PEERS cmd

func generateREQ_PEERS() string {
	return "REQ_PEERS\n"
}

func handleREQ_PEERS(cmd []string, client tcp.TCPConnection) error {
	var ps []struct {
		hostname string
		port     uint
	}

	p := Peer()
	p.peersMU.Lock()
	clientPD, ok := p.peers[client]
	if !ok {
		return nil
	}

	for _, pd := range p.peers {
		if pd.Hostname == clientPD.Hostname && pd.Port == clientPD.Port {
			continue
		}

		if p.server.IsClient(client) {
			continue
		}

		ps = append(ps, struct {
			hostname string
			port     uint
		}{pd.Hostname, pd.Port})
	}
	p.peersMU.Unlock()

	client.SendString(generatePEERS(ps))
	return nil
}

// PEERS cmd

func generatePEERS(peers []struct {
	hostname string
	port     uint
}) string {
	if len(peers) == 0 {
		return "PEERS\n"
	}

	var sb strings.Builder
	sb.WriteString("PEERS")
	for i := 0; i < len(peers); i++ {
		sb.WriteString(fmt.Sprintf(";%s:%d", peers[i].hostname, peers[i].port))
	}
	sb.WriteString("\n")

	return sb.String()

}

func handlePEERS(cmd []string, client tcp.TCPConnection) error {
	p := Peer()
	counter := 0

	var pds []*peerData
	p.peersMU.Lock()
	for _, pd := range p.GetPeers() {
		pds = append(pds, pd)
	}
	p.peersMU.Unlock()

	myHostname, myPort := p.GetServerInfo()

	for i := 1; i < len(cmd); i++ {
		parts := strings.Split(cmd[i], ":")
		port, err := strconv.ParseUint(parts[1], 10, 32)
		if err != nil {
			continue
		}

		if parts[0] == myHostname && uint(port) == myPort {
			continue
		}

		for _, pd := range pds {
			if !(pd.Hostname == parts[0] && pd.Port == uint(port)) {
				if err := p.ConnectToPeer(parts[0], uint(port)); err != nil {
					util.Logger.Warnf("err connecting to peer %s:%d: %s", parts[0], port, err)
				} else {
					counter++
				}
				break
			}
		}
	}

	util.Logger.Infof("Attempting to connect to %d new peers", counter)
	return nil
}
