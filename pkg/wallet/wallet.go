package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"

	"github.com/SIGBlockchain/aurum_client/pkg/constants"
	"github.com/SIGBlockchain/aurum_client/pkg/privatekey"
)

func SetupWallet() error {
	// if the JSON file already exists, return error
	_, err := os.Stat(constants.Wallet)
	if err == nil {
		return errors.New("JSON file for aurum_wallet already exists")
	}

	// Create JSON file for wallet
	file, err := os.Create(constants.Wallet)
	if err != nil {
		return err
	}
	defer file.Close()

	// Json structure that will be used to store information into the json file
	type jsonStruct struct {
		PrivateKey string
		Balance    uint64
		Nonce      uint64
	}

	// Generate ecdsa key pairs
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return err
	}

	// Encodes private key
	pemEncoded, err := privatekey.EncodePrivateKey(privateKey)
	if err != nil {
		return err
	}

	// Encodes the pem encoded private key into string and stores it into a jsonStruct
	hexKey := hex.EncodeToString(pemEncoded)
	j := jsonStruct{PrivateKey: hexKey}

	// Marshall the jsonStruct
	jsonEncoded, err := json.Marshal(j)
	if err != nil {
		return err
	}

	// Write into the json file
	_, err = file.Write(jsonEncoded)
	if err != nil {
		return err
	}

	return nil
}
