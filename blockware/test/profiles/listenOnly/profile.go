package profileListenOnly

import (
	"math/big"
	"os"
	"time"

	"github.com/t02smith/part-iii-project/toolkit/model/games"
	"github.com/t02smith/part-iii-project/toolkit/model/net"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*

Profile 1:

Features:
- listen-and-respond only peer
- will upload a game to ETH and seed it
- ideal connection

Data:

- small size project (latex-template from github)
	a small scale project that will allow for the peer to download
	from many different files.

- fake large file
	TODO ^^
	A large amount of blocks that can be easily generated without
	storing them => e.g. block 0 has a content of ...0000, block 1
	is ...00001, etc.
	Will only need to store the hash with the offset to be able to
	generate it

*/

var PrivateKey string = testutil.Accounts[2][1]

func Run() {
	// _, err := generateLargeFile()
	// if err != nil {
	// 	util.Logger.Fatalf("Error generating large file: %s", err)
	// }

	// ? PRE-LAUNCH CHECKS
	_, err := os.Stat("../listenOnlyWithUpload/latex-template-main")
	if err != nil {
		util.Logger.Fatalf("Latex template directory not found. Run 'make' to fetch it")
	}

	p := net.Peer()

	// * create & upload game
	g, err := games.CreateGame(games.NewGame{
		Title:       "latex-template",
		Version:     "4.7.1",
		ReleaseDate: time.Date(2002, time.January, 10, 0, 0, 0, 0, time.UTC).String(),
		Developer:   "tcs1g20",
		RootDir:     "../listenOnlyWithUpload/latex-template-main",
		Price:       big.NewInt(150),
		ShardSize:   1024},
		nil,
	)

	if err != nil {
		util.Logger.Fatalf("Error creating game: %s", err)
	}

	p.Library().AddOwnedGame(g)
	err = games.OutputAllGameDataToFile(g)
	if err != nil {
		util.Logger.Warnf("Error saving game to file %s", err)
	}

	// * listen for connections
	util.Logger.Info("PROFILE 1: listening")
	<-make(chan bool) // wait forever :/
}
