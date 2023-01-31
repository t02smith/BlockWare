/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/t02smith/part-iii-project/toolkit/cmd"
	"github.com/t02smith/part-iii-project/toolkit/lib"
)

func main() {
	SetupConfig()
	lib.InitLogger()
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	lib.SetupToolkitEnvironment()
	cmd.Execute()
}
