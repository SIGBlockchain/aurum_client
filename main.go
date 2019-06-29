package main

import (
	"fmt"
	"os"

	"github.com/SIGBlockchain/aurum_client/pkg/wallet"
	"github.com/pborman/getopt"
)

type Opts struct {
	help       *bool
	version    *bool
	setup      *bool
	info       *bool
	updateInfo *bool
	contract   *bool
	recipient  *string
	value      *string
	producer   *string
}

// TODO: Use flag package
func main() {
	options := Opts{
		help:       getopt.BoolLong("help", '?', "help"),
		version:    getopt.BoolLong("version", 'w', "version"),
		setup:      getopt.BoolLong("setup", 's', "set up client"),
		info:       getopt.BoolLong("info", 'i', "wallet info"),
		updateInfo: getopt.BoolLong("update", 'u', "update wallet info"),
		contract:   getopt.BoolLong("contract", 'c', "make contract"),
		recipient:  getopt.StringLong("recipient", 'r', "recipient"),
		value:      getopt.StringLong("value", 'v', "", "value to send"),
		producer:   getopt.StringLong("producer", 'p', "", "producer address"),
	}

	if *options.help {
		getopt.Usage()
		os.Exit(0)
	}

	if *options.setup {
		fmt.Println("Initializing Aurum wallet...")
		if err := wallet.SetupWallet(); err != nil {
			fmt.Printf("Failed to set up wallet: %v", err)
			os.Exit(1)
		}
		fmt.Println("Wallet setup complete.")
		if err := wallet.PrintInfo(); err != nil {
			fmt.Printf("Failed to print wallet info: %v", err)
		}
	}
}
