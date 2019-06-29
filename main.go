package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/SIGBlockchain/aurum_client/pkg/wallet"
)

type Opts struct {
	setup      *bool
	info       *bool
	updateInfo *bool
	value      *string
	recipient  *string
}

// TODO: Use flag package
func main() {
	options := Opts{
		setup:      flag.Bool("setup", false, "set up client"),
		info:       flag.Bool("info", false, "wallet info"),
		updateInfo: flag.Bool("update", false, "update wallet info"),
		recipient:  flag.String("to", "", "recipient"),
		value:      flag.String("send", "", "value to send"),
	}
	flag.Parse()

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
