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
	verifyDirToVerify      string
	verifyExpectedHashFile string
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		f, err := lib.ReadHashTreeFromFile(verifyExpectedHashFile)
		if err != nil {
			fmt.Printf("error reading hash file: %s\n", err)
		}
		fmt.Println(f)
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)

	verifyCmd.Flags().StringVarP(&verifyDirToVerify, "directory", "d", "", "Which directory are you attempting to verify")
	verifyCmd.Flags().StringVarP(&verifyExpectedHashFile, "hash-file", "f", "hash.json", "The location of your hash file with your expected results")
}
