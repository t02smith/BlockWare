package net

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/lib/games"
)

func onMessage(cmd []string, client PeerIT) {
	p := GetPeerInstance()

	switch cmd[0] {

	// LIBRARY => request a list of a peers games
	case "LIBRARY":
		log.Println("Library command called")

		gameLs, err := games.LoadGames(viper.GetString("meta.directory"))
		if err != nil {
			log.Printf("Error loading games: %s\n", err)
			return
		}

		gameStr, err := gameListToMessage(gameLs)
		if err != nil {
			log.Printf("Error serialising games: %s\n", err)
			return
		}

		client.SendString(gameStr)
		return

	// GAMES => a list of users games
	case "GAMES":
		log.Println("Games command called")

		ls, err := gameMessageToGameList(cmd)
		if err != nil {
			log.Printf("Error reading games: %s\n", err)
			return
		}

		if peer, ok := p.peers[client]; ok {
			log.Println("Client found. Updating library")
			peer.Library = ls
		}

		return

	// BLOCK <hash> => Request a block of data from a user
	case "BLOCK":
		log.Printf("Block command called for block %s", cmd[1])

		// haveBlock, err :=

		return

	// SEND_BLOCK <game hash> <hash> <shard> => Download a shard off of a user
	case "SEND_BLOCK":
		return

	}
}

//

// turns a game library to a message to send to a peer
func gameListToMessage(games []*games.Game) (string, error) {
	var buf bytes.Buffer
	buf.WriteString("GAMES;")

	for _, g := range games {
		encoded, err := g.Serialise()
		if err != nil {
			return "", nil
		}

		buf.WriteString(fmt.Sprintf("%s;", encoded))
	}

	buf.WriteString("\n")
	return buf.String(), nil
}

// takes a game library message and converts it to a list of games
func gameMessageToGameList(parts []string) ([]*games.Game, error) {
	gameLs := []*games.Game{}

	for i := 1; i < len(parts); i++ {
		g, err := games.DeserialiseGame(parts[i])

		if err == io.EOF {
			return gameLs, nil
		}

		if err != nil {
			return nil, err
		}

		gameLs = append(gameLs, g)
	}

	return gameLs, nil
}

// fetch a block given a game identifier and a shard
func fetchBlock(gameHash, shardHash [32]byte) ([]byte, error) {
	p := GetPeerInstance()

	// find the game
	g, d := p.library.GetGame(gameHash)
	if g == nil {
		return nil, errors.New("game not in library")
	}

	// game is being downloaded and we may not have block
	if d != nil {
		if _, ok := d.Progress[shardHash]; ok {
			return nil, errors.New("shard hasn't been downloaded yet")
		}
	}

	// fetch the block
	b, err := g.FetchShard(shardHash)
	if err != nil {
		return nil, err
	}

	return b, nil
}
