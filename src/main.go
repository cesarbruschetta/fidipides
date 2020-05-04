package main

import (
	"./cmd"
	"./cmd/utils"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		utils.PrintErrorAndExit(err.Error())
	}
}
