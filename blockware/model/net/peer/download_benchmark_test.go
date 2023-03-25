package peer

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/t02smith/part-iii-project/toolkit/model/manager/games"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*

This benchmark will test the performance of downloading a game
over blockchain. It will consider three different variables:
1. number of peers
2. number of files
3. game size

Assumptions for all benchmarks:
- the same system is used for all benchmarks

*/

const (
	B_RAND_SEED int64 = 42
)

/*

Benchmark: Peer Count P

assumptions:
- the same game is used between benchmarks
- all files are in the root directory

P = [1, 2, 4, 8, 16] peers

*/

func BenchmarkDownloadGame_PeerCount(b *testing.B) {
	b.Skip()

	// 500 * 80MB files = 40GB game
	err := testutil.GenerateLargeFolder("peer-count", "../../../test/data/tmp/games", 80_000_000, 500)
	if err != nil {
		b.Fatal(err)
	}

	_, err = games.CreateGame(games.NewGame{
		Title:       "peer-count",
		Version:     "4.7.1",
		ReleaseDate: time.Date(2002, time.January, 10, 0, 0, 0, 0, time.UTC).String(),
		Developer:   "tcs1g20",
		RootDir:     "../../../test/data/tmp/games/peer-count",
		Price:       big.NewInt(150),
		ShardSize:   4194304,
		AssetsDir:   "../../../test/data/assets"},
		nil,
	)
	if err != nil {
		b.Fatal(err)
	}

	peerCounts := []int{1}
	for _, peerCount := range peerCounts {
		util.Logger.Infof("BENCHMARK - START - Peer Count - %d", peerCount)

		b.Run(fmt.Sprintf("Peer count = %d", peerCount), func(b *testing.B) {
			for i := 0; i < peerCount; i++ {

			}
		})

		util.Logger.Infof("BENCHMARK - END - Peer Count - %d", peerCount)
	}

}
