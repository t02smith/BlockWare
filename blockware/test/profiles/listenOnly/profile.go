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

*/

var PrivateKey string = testutil.Accounts[2][1]

func Run() {
	// ? PRE-LAUNCH CHECKS
	if _, err := os.Stat("../../data/tmp/games/profile-2"); err != nil {
		if err = testutil.GenerateLargeFolder("profile-2", "../../data/tmp/games/", 80_000_000, 500); err != nil {
			util.Logger.Fatal(err)
		}
	}

	p := peer.Peer()

	// * create & upload game
	// g, err := games.CreateGame(games.NewGame{
	// 	Title:       "t02smith.github.io",
	// 	Version:     "4.7.1",
	// 	ReleaseDate: time.Date(2002, time.January, 10, 0, 0, 0, 0, time.UTC).String(),
	// 	Developer:   "tcs1g20",
	// 	RootDir:     "../listenOnlyWithUpload/t02smith.github.io",
	// 	Price:       big.NewInt(150),
	// 	ShardSize:   4194304,
	// 	AssetsDir:   "../../data/assets"},
	// 	nil,
	// )

	g, err := games.CreateGame(games.NewGame{
		Title:       "benchmark",
		Version:     "4.7.1",
		ReleaseDate: time.Date(2002, time.January, 10, 0, 0, 0, 0, time.UTC).String(),
		Developer:   "tcs1g20",
		RootDir:     "../../data/tmp/games/profile-2",
		Price:       big.NewInt(150),
		ShardSize:   67108864,
		AssetsDir:   "../../data/assets"},
		nil,
	)

	if err != nil {
		util.Logger.Fatalf("Error creating game: %s", err)
	}

	p.Library().AddOrUpdateOwnedGame(g)
	err = games.OutputAllGameDataToFile(g)
	if err != nil {
		util.Logger.Warnf("Error saving game to file %s", err)
	}

	// * listen for connections
	util.Logger.Info("PROFILE 2: listening")
	<-make(chan bool) // wait forever :/
}
