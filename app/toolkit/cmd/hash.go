/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/lib/io"
)

var (
	hashDirectory string
	hashShardSize int
	workerCount   int
)

// hashCmd represents the hash command
var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Generate the hash tree for a game directory.",
	Long: `This function will take the root directory of the video
game you want to upload and generate a hash tree containing the 
hashes of each shard of data. This will allow your users to 
verify the contents they are downloading.`,
	Run: func(cmd *cobra.Command, args []string) {
		t, err := io.NewHashTree(hashDirectory, uint(hashShardSize))
		if err != nil {
			fmt.Println(err)
		}

		viper.Set("meta.hashes.workerCount", workerCount)

		t.Hash()
		t.OutputToFile(fmt.Sprintf(".%x.json", t.RootDir.RootHash))
	},
}

func init() {
	rootCmd.AddCommand(hashCmd)

	hashCmd.Flags().StringVarP(&hashDirectory, "directory", "d", "", "The path to the directory you'd like to hash")
	hashCmd.Flags().IntVarP(&hashShardSize, "shard-size", "s", 16384, "The size (in bytes) of each shard to be hashed")
	hashCmd.Flags().IntVarP(&workerCount, "worker-count", "w", 5, "How many worker threads to hash with")
}
