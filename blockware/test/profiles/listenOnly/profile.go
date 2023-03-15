package profileListenOnly

import (
	"math/big"
	"os"
	"time"

	"github.com/t02smith/part-iii-project/toolkit/model/manager/games"
	"github.com/t02smith/part-iii-project/toolkit/model/net/peer"
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
	_, err := os.Stat("../listenOnlyWithUpload/t02smith.github.io")
	if err != nil {
		util.Logger.Fatalf("Latex template directory not found. Run 'make' to fetch it")
	}

	p := peer.Peer()

	// * create & upload game
	g, err := games.CreateGame(games.NewGame{
		Title:       "t02smith.github.io",
		Version:     "4.7.1",
		ReleaseDate: time.Date(2002, time.January, 10, 0, 0, 0, 0, time.UTC).String(),
		Developer:   "tcs1g20",
		RootDir:     "../listenOnlyWithUpload/t02smith.github.io",
		Price:       big.NewInt(150),
		ShardSize:   16384,
		AssetsDir:   "../../data/assets"},
		nil,
	)

	// g, err := games.CreateGame(games.NewGame{
	// 	Title:       "Transformers",
	// 	Version:     "4.7.1",
	// 	ReleaseDate: time.Date(2002, time.January, 10, 0, 0, 0, 0, time.UTC).String(),
	// 	Developer:   "tcs1g20",
	// 	RootDir:     "C:\\Games\\Transformers - Fall of Cybertron",
	// 	Price:       big.NewInt(150),
	// 	ShardSize:   4194304,
	// 	AssetsDir:   "../../data/assets"},
	// 	nil,
	// )

	if err != nil {
		util.Logger.Fatalf("Error creating game: %s", err)
	}

	p.Library().AddOwnedGame(g)
	err = games.OutputAllGameDataToFile(g)
	if err != nil {
		util.Logger.Warnf("Error saving game to file %s", err)
	}

	// * listen for connections
	util.Logger.Info("PROFILE 2: listening")
	<-make(chan bool) // wait forever :/
}
