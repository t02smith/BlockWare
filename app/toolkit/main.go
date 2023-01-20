/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"
	"os"

	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/cmd"
)

func main() {
	SetupConfig()

	logLocation := viper.GetString("meta.log")
	if logLocation != "stdout" {
		f, err := os.OpenFile(logLocation, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Error opening log file %s: %s", logLocation, err)
		}
		defer f.Close()

		log.SetOutput(f)
	}

	cmd.Execute()
}
