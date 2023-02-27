package hash

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"testing"

	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*

Hash Tree Benchmark
This benchmark will test how the HashTree.Hash function performs when we vary
the following variables:
1. file count
2. avg. file size
3. worker count
4. shard size

*/

const (
	B_RAND_SEED int64 = 42
)

/*

Benchmark: file count

assumptions:
- all files are in the same directory
- all files are the same size
- all files contain random data
- same worker count

N = [100, 1000, 10000, 100000] files

TODO fetch data from online

*/

const (
	b_filecount_filesize  uint = 4096 // bytes
	b_filecount_shardsize uint = 1024 // bytes
	b_filecount_workers   uint = 5
)

func _generateFileCountData(n uint) error {
	err := os.Mkdir("../../test/data/tmp/benchmark_filecount", 0644)
	if err != nil {
		return err
	}

	rand.Seed(B_RAND_SEED)
	data := make([]byte, b_filecount_filesize)
	for j := 0; j < int(b_filecount_filesize); j++ {
		data[j] = byte(rand.Intn(256))
	}

	for i := 0; i < int(n); i++ {
		f, err := os.Create(filepath.Join("../../test/data/tmp/benchmark_filecount", fmt.Sprint(i)))
		if err != nil {
			return err
		}

		writer := bufio.NewWriter(f)
		_, err = writer.Write(data)
		if err != nil {
			return err
		}

		err = writer.Flush()
		if err != nil {
			return err
		}

		f.Close()
	}

	return nil
}

func BenchmarkHash_FileCount(b *testing.B) {
	util.InitLogger()

	levels := []uint{100, 1_000, 10_000, 100_000}
	for _, n := range levels {
		util.Logger.Infof("Generating data for file count = %d", n)
		err := _generateFileCountData(n)
		if err != nil {
			b.Fatal(err)
		}
		util.Logger.Infof("Generated data for file count = %d", n)

		b.Run(fmt.Sprintf("file count = %d", n), func(b *testing.B) {

			ht, err := NewHashTree("../../test/data/tmp/benchmark_filecount", b_filecount_shardsize, nil)
			if err != nil {
				b.Fatal(err)
			}

			util.Logger.Infof("Starting hash for file count = %d", n)
			err = ht.Hash()
			if err != nil {
				b.Fatal(err)
			}
			util.Logger.Infof("Finished hash for file count = %d", n)
		})

		testutil.ClearTmp("../../")
	}
}

/*

Benchmark: file size

assumptions:
- there are always the same number of files
- all files are the same size
- all files contain random data
- same worker count

*/

const (
	b_filesize_filecount uint = 250
	b_filesize_shardsize uint = 1024 // bytes
	b_filesize_workers   uint = 5
)

// generate the data directories
func _generateFileSizeData(size uint) error {
	return nil
}

func BenchmarkHash_FileSize(b *testing.B) {

}

/*

Benchmark: worker count

assumptions:
- constant file count
- constant file size
- all files contain random data

*/

const (
	b_workercount_filesize  uint = 4096
	b_workercount_shardsize uint = 1024
	b_workercount_filecount uint = 250
)

func _generateWorkerCountData() error {
	return nil
}

func BenchmarkHash_WorkerCount(b *testing.B) {

}

/*

Benchmark: shard size

assumptions:
- constant file size
- constant file count
- constant worker count

*/

const (
	b_shardsize_filesize    uint = 4096
	b_shardsize_filecount   uint = 250
	b_shardsize_workercount uint = 5
)

func _generateShardSizeData() error {
	return nil
}

func BenchmarkHash_ShardSize(b *testing.B) {

}
