package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/SIGBlockchain/aurum_client/internal/config"
	"github.com/SIGBlockchain/aurum_client/internal/contracts"
	"github.com/SIGBlockchain/aurum_client/internal/requests"
	"github.com/SIGBlockchain/aurum_client/internal/wallet"
)

type Opts struct {
	version   *bool
	setup     *bool
	info      *bool
	update    *bool
	value     *string
	recipient *string
	producer  *string
	save      *string
	as        *string
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
		producer:  flag.String("producer", cfg.ProducerAddress, "IP address of the producer"),
		save:      flag.String("save", "", "wallet addresss"),
		as:        flag.String("as", "", "name"),
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
		return
	}

	if *options.producer != cfg.ProducerAddress {
		//check if a valid ip was given
		parts := strings.Split(*options.producer, ":")
		if len(parts) != 2 {
			log.Fatalf("%s is not a valid ip:port address", *options.producer)
		}
		port, err := strconv.Atoi(parts[1])
		if err != nil || port > 65535 {
			log.Fatalf("Invalid port %s", parts[1])
		}
		if net.ParseIP(parts[0]) == nil {
			log.Fatalf("Invalid IP address \"%s\"", parts[0])
		}
		cfg.ProducerAddress = *options.producer
		defer config.SaveConfiguration(cfg)
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
		return
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
		if resp.StatusCode != http.StatusOK {
			if resp.StatusCode == http.StatusNotFound { //account requested was no valid
				bodyBytes, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatal(err)
				}
				bodyString := string(bodyBytes)
				log.Fatalf("Account status could not be retrievd: %v\n", bodyString)
			}
			log.Fatalf("Status code: %v\n", resp.StatusCode)
		}
		defer resp.Body.Close()
		updatedWallet := new(wallet.Wallet)
		body, err := ioutil.ReadAll(resp.Body)
		if err := json.Unmarshal(body, &updatedWallet); err != nil {
			log.Fatalf("Failed to unmarshall response body: %v\n", err)
		}
		defer resp.Body.Close()
		if err := wallet.UpdateWallet(updatedWallet.Balance, updatedWallet.StateNonce); err != nil {
			log.Fatalf("Failed to update wallet: %v\n", err)
		}
		log.Println("Wallet updated successfully.")
		return
	}

	if *options.value != "" && *options.recipient != "" {
		if !wallet.ValidRecipLen(*options.recipient) {
			log.Fatalf("Recipient input is not 64 characters long\n")
		}
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
			// TODO: Try again with incremented nonce
			log.Println("Please wait for next block to be produced or update wallet.")
			log.Fatalf("Status code: %v\nBody: %s\n", resp.StatusCode, buf.String())
		}
		defer resp.Body.Close()
		currentBalance, err := wallet.GetBalance()
		if err != nil {
			log.Fatalf("Failed to get current balance: %v\n", err)
		}
		currentNonce, err := wallet.GetStateNonce()
		if err != nil {
			log.Fatalf("Failed to get current nonce: %v\n", err)
		}
		intVal, err := strconv.Atoi(*options.value)
		if err != nil {
			log.Fatalf("Failed to convert value to integer: %v", err)
		}
		if err := wallet.UpdateWallet(currentBalance-uint64(intVal), currentNonce+1); err != nil {
			log.Fatalf("Failed to update wallet: %v\n", err)
		}
		log.Println("Successfully sent contract to producer.\nContract will be confirmed once next block is produced.")
	}

	if *options.save != "" && *options.as != "" {
		if !wallet.ValidRecipLen(*options.save) {
			log.Fatalf("Recipient input is not 64 characters long\n")
		}

		// open addressbook json to add wallet address

	}

}

func checkFlagCount(limit int) {
	if flag.NFlag() > limit {
		log.Fatalln("Too many arguments")
	}
}
