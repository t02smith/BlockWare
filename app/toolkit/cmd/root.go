/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/t02smith/part-iii-project/toolkit/cmd/new"
	"github.com/t02smith/part-iii-project/toolkit/cmd/start"
	"github.com/t02smith/part-iii-project/toolkit/cmd/view"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "toolkit",
	Short: "A CLI tool to generate data and verify downloaded files",
	Long: `This CLI tool will allow you to quickly generate and verify
	the data required to interact with my blockchain video game marketplace`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(start.StartCmd)
	rootCmd.AddCommand(new.NewCmd)
	rootCmd.AddCommand(view.ViewCmd)
}
