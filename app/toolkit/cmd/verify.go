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
	verifyDirToVerify           string
	verifyExpectedHashFile      string
	verifyIgnoreNewFilesAndDirs bool
	verifyContinueAfterError    bool
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
			return
		}

		config := &lib.VerifyHashTreeConfig{
			IgnoreNewFilesAndDirs: verifyIgnoreNewFilesAndDirs,
			ContinueAfterError:    verifyContinueAfterError,
		}

		res, err := f.VerifyTree(config, verifyDirToVerify)
		if err != nil {
			fmt.Printf("error verifying directory: %s\n", err)
			return
		}

		fmt.Printf("Directory matches: %t\n", res)
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)

	verifyCmd.Flags().StringVarP(&verifyDirToVerify, "directory", "d", "", "Which directory are you attempting to verify")
	verifyCmd.Flags().StringVarP(&verifyExpectedHashFile, "hash-file", "f", "hash.json", "The location of your hash file with your expected results")

	verifyCmd.Flags().BoolVarP(&verifyIgnoreNewFilesAndDirs, "ignore-new-data", "i", false, "Ignore new files or directories that have been added")
	verifyCmd.Flags().BoolVarP(&verifyContinueAfterError, "continue-after-error", "c", false, "Don't stop the verification if an error is found")
}
