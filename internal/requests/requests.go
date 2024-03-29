package requests

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/SIGBlockchain/aurum_client/internal/contracts"
	"github.com/SIGBlockchain/aurum_client/internal/endpoints"
	"github.com/SIGBlockchain/aurum_client/internal/publickey"
)

type JSONContract struct {
	Version                uint16
	SenderPublicKey        string
	SignatureLength        uint8
	Signature              string
	RecipientWalletAddress string
	Value                  uint64
	StateNonce             uint64
}

func NewAccountInfoRequest(host string, walletAddress string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, "http://"+host+endpoints.AccountInfo, nil)
	if err != nil {
		return nil, errors.New("Failed to make new request:\n" + err.Error())
	}
	values := req.URL.Query()
	values.Add("w", walletAddress)
	req.URL.RawQuery = values.Encode()
	return req, nil
}

func NewContractRequest(host string, newContract contracts.Contract) (*http.Request, error) {
	// TODO: accounts.Contract to JSON Call it MarshalContract?
	var newJSONContract = JSONContract{
		Version:                newContract.Version,
		SenderPublicKey:        hex.EncodeToString(publickey.EncodePublicKey(newContract.SenderPubKey)),
		SignatureLength:        newContract.SigLen,
		Signature:              hex.EncodeToString(newContract.Signature),
		RecipientWalletAddress: hex.EncodeToString(newContract.RecipPubKeyHash),
		Value:      newContract.Value,
		StateNonce: newContract.StateNonce,
	}
	marshalledContract, err := json.Marshal(newJSONContract)
	if err != nil {
		return nil, errors.New("Failed to marshall contract: " + err.Error())
	}
	req, err := http.NewRequest(http.MethodPost, "http://"+host+endpoints.Contract, bytes.NewBuffer(marshalledContract))
	if err != nil {
		return nil, errors.New("Failed to create request: " + err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil

}
