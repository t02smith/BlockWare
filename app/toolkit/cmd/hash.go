/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/t02smith/part-iii-project/toolkit/lib"
)

var (
	hashDirectory string
	hashShardSize uint
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
		t, err := lib.NewHashTree(hashDirectory, hashShardSize)
		if err != nil {
			fmt.Println(err)
		}

		t.Hash()
		t.OutputToFile(fmt.Sprintf(".%x.json", t.RootDir.RootHash))
	},
}

func init() {
	rootCmd.AddCommand(hashCmd)

	hashCmd.Flags().StringVarP(&hashDirectory, "directory", "d", "", "The path to the directory you'd like to hash")
	hashCmd.Flags().UintVarP(&hashShardSize, "shard-size", "s", 1024, "The size (in bytes) of each shard to be hashed")
}
