package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/SIGBlockchain/aurum_client/pkg/contracts"

	"github.com/SIGBlockchain/aurum_client/pkg/config"
	"github.com/SIGBlockchain/aurum_client/pkg/requests"
	"github.com/SIGBlockchain/aurum_client/pkg/wallet"
)

type Opts struct {
	version   *bool
	setup     *bool
	info      *bool
	update    *bool
	value     *string
	recipient *string
}

func main() {
	log.SetFlags(0)
	cfg, err := config.LoadConfiguration()
	if err != nil {
		log.Printf("Failed to load configuration file: %v", err)
		os.Exit(1)
	}
	options := Opts{
		version:   flag.Bool("version", false, "client version"),
		setup:     flag.Bool("setup", false, "set up client"),
		info:      flag.Bool("info", false, "wallet info"),
		update:    flag.Bool("update", false, "update wallet info"),
		recipient: flag.String("to", "", "recipient"),
		value:     flag.String("send", "", "value to send"),
	}
	flag.Parse()

	if *options.version {
		checkFlagCount(1)
		log.Printf("Aurum client version: %d\n", cfg.Version)
	}

	if *options.info {
		checkFlagCount(1)
		if err := wallet.PrintInfo(); err != nil {
			log.Fatalf("Failed to get wallet contents: %v\n", err)
		}
		os.Exit(0)
	}

	if *options.setup {
		checkFlagCount(1)
		log.Println("Initializing Aurum wallet...")
		if err := wallet.SetupWallet(); err != nil {
			log.Fatalf("Failed to set up wallet: %v\n", err)
		}
		log.Println("Wallet setup complete.")
		if err := wallet.PrintInfo(); err != nil {
			log.Fatalf("Failed to print wallet info: %v\n", err)
		}
		os.Exit(0)
	}

	if *options.update {
		checkFlagCount(1)
		cli := new(http.Client)
		walletAddress, err := wallet.GetWalletAddress()
		if err != nil {
			log.Fatalf("Failed to extract wallet address: %v\n", err)
		}
		req, err := requests.NewAccountInfoRequest(cfg.ProducerAddress, hex.EncodeToString(walletAddress))
		if err != nil {
			log.Fatalf("Failed to make update request: %v\n", err)
		}
		fmt.Println("Requesting wallet update from producer...")
		resp, err := cli.Do(req)
		if err != nil {
			log.Fatalf("Failed getting response from producer: %v\n", err)
		}
		if resp.StatusCode != http.StatusFound {
			log.Fatalf("Status code: %v\n", resp.StatusCode)
		}
		defer resp.Body.Close()
		updatedWallet := new(wallet.Wallet)
		body, err := ioutil.ReadAll(resp.Body)
		if err := json.Unmarshal(body, &updatedWallet); err != nil {
			log.Fatalf("Failed to unmarshall response body: %v\n", err)
		}
		if err := wallet.UpdateWallet(updatedWallet.Balance, updatedWallet.StateNonce); err != nil {
			log.Fatalf("Failed to update wallet: %v\n", err)
		}
		log.Println("Wallet updated successfully.")
		os.Exit(0)
	}

	if *options.value != "" && *options.recipient != "" {
		newContract, err := contracts.ContractMessageFromInput(cfg.Version, *options.value, *options.recipient)
		if err != nil {
			log.Fatalf("Failed to construct new contract: %v\n", err)
		}
		cli := new(http.Client)
		req, err := requests.NewContractRequest(cfg.ProducerAddress, *newContract)
		if err != nil {
			log.Fatalf("Failed to make contract request: %v\n", err)
		}
		fmt.Printf("Sending contract request with %s to %s", *options.value, *options.recipient)
		resp, err := cli.Do(req)
		if err != nil {
			log.Fatalf("Failed getting response from producer: %v\n", err)
		}
		if resp.StatusCode != http.StatusOK {
			buf := new(bytes.Buffer)
			_, err := resp.Body.Read(buf.Bytes())
			if err != nil {
				log.Fatalf("Failed to read body of response: %v\n", err)
			}
			log.Fatalf("Status code: %v\nBody: %s\n", resp.StatusCode, buf.String())
		}
		defer resp.Body.Close()
		fmt.Println("Successfully sent contract to producer.\nContract will be confirmed once next block is producer.")
	}

}

func checkFlagCount(limit int) {
	if flag.NFlag() > limit {
		log.Fatalln("Too many arguments")
	}
}
