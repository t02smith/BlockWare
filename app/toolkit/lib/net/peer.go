package net

import (
	"bytes"
	"fmt"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/lib/games"
)

type Peer struct {

	// connections
	server  *TCPServer
	clients []*TCPClient

	// data
	installFolder string
	games         []*games.Game
}

func StartPeer(serverHostname string, serverPort uint, installFolder, gameDataLocation string) error {
	gameLs, err := games.LoadGames(gameDataLocation)
	if err != nil {
		return err
	}

	fmt.Printf("Found %d games\n", len(gameLs))
	p := &Peer{
		server:        InitServer(serverHostname, serverPort),
		clients:       []*TCPClient{},
		installFolder: installFolder,
		games:         gameLs,
	}

	p.server.Start(onMessage)
	return nil
}

//

func onMessage(cmd []string, client *TCPServerClient) {
	switch cmd[0] {

	// LIBRARY => request a list of a peers games
	case "LIBRARY":
		fmt.Println("Library command called")

		gameLs, err := games.LoadGames(viper.GetString("meta.directory"))
		if err != nil {
			fmt.Printf("Error loading games: %s\n", err)
			return
		}

		gameStr, err := gameListToMessage(gameLs)
		if err != nil {
			fmt.Printf("Error serialising games: %s\n", err)
			return
		}

		client.send(gameStr)
		return

	// GAMES => a list of users games
	case "GAMES":
		fmt.Println("Games command called")

		_, err := gameMessageToGameList(cmd)
		if err != nil {
			fmt.Printf("Error reading games: %s\n", err)
			return
		}

		return

	}
}

//

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

func gameMessageToGameList(parts []string) ([]*games.Game, error) {
	gameLs := []*games.Game{}

	for i := 1; i < len(parts); i++ {
		g, err := games.DeserialiseGame(parts[i])
		if err != nil {
			return nil, err
		}

		gameLs = append(gameLs, g)
	}

	return gameLs, nil
}
