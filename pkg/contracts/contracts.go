package contracts

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"strconv"

	"github.com/SIGBlockchain/aurum_client/pkg/hashing"
	"github.com/SIGBlockchain/aurum_client/pkg/publickey"
	"github.com/SIGBlockchain/aurum_client/pkg/wallet"
)

type Contract struct {
	Version         uint16
	SenderPubKey    *ecdsa.PublicKey
	SigLen          uint8
	Signature       []byte
	RecipPubKeyHash []byte
	Value           uint64
	StateNonce      uint64
}

func MakeContract(version uint16, sender *ecdsa.PrivateKey, recipient []byte, value uint64, nextStateNonce uint64) (*Contract, error) {

	if version == 0 {
		return nil, errors.New("Invalid version; must be >= 1")
	}

	c := Contract{
		Version:         version,
		SigLen:          0,
		Signature:       nil,
		RecipPubKeyHash: recipient,
		Value:           value,
		StateNonce:      nextStateNonce,
	}

	if sender == nil {
		c.SenderPubKey = nil
	} else {
		c.SenderPubKey = &(sender.PublicKey)
	}

	return &c, nil
}

// // Serialize all fields of the contract
func (c *Contract) Serialize() ([]byte, error) {
	/*
		0-2 version
		2-180 spubkey
		180-181 siglen
		181 - 181+c.siglen signature
		181+c.siglen - (181+c.siglen + 32) rpkh
		(181+c.siglen + 32) - (181+c.siglen + 32+ 8) value

	*/

	// if contract's sender pubkey is nil, make 178 zeros in its place instead
	var spubkey []byte
	if c.SenderPubKey == nil {
		spubkey = make([]byte, 178)
	} else {
		spubkey = publickey.EncodePublicKey(c.SenderPubKey) //size 178
	}

	//unsigned contract
	if c.SigLen == 0 {
		totalSize := (2 + 178 + 1 + 32 + 8 + 8)
		serializedContract := make([]byte, totalSize)
		binary.LittleEndian.PutUint16(serializedContract[0:2], c.Version)
		copy(serializedContract[2:180], spubkey)
		serializedContract[180] = 0
		copy(serializedContract[181:213], c.RecipPubKeyHash)
		binary.LittleEndian.PutUint64(serializedContract[213:221], c.Value)
		binary.LittleEndian.PutUint64(serializedContract[221:229], c.StateNonce)

		return serializedContract, nil
	} else { //signed contract
		totalSize := (2 + 178 + 1 + int(c.SigLen) + 32 + 8 + 8)
		serializedContract := make([]byte, totalSize)
		binary.LittleEndian.PutUint16(serializedContract[0:2], c.Version)
		copy(serializedContract[2:180], spubkey)
		serializedContract[180] = c.SigLen
		copy(serializedContract[181:(181+int(c.SigLen))], c.Signature)
		copy(serializedContract[(181+int(c.SigLen)):(181+int(c.SigLen)+32)], c.RecipPubKeyHash)
		binary.LittleEndian.PutUint64(serializedContract[(181+int(c.SigLen)+32):(181+int(c.SigLen)+32+8)], c.Value)
		binary.LittleEndian.PutUint64(serializedContract[(181+int(c.SigLen)+32+8):(181+int(c.SigLen)+32+8+8)], c.StateNonce)

		return serializedContract, nil
	}
}

func (c *Contract) SignContract(sender *ecdsa.PrivateKey) error {
	serializedTestContract, err := c.Serialize()
	if err != nil {
		return errors.New("Failed to serialize contract")
	}
	hashedContract := hashing.HashSHA256(serializedTestContract)
	c.Signature, _ = sender.Sign(rand.Reader, hashedContract, nil)
	c.SigLen = uint8(len(c.Signature))
	return nil
}

// Convert value to uint64; if unsuccessful output an error
// If value is zero, output error
// GetBalance(), if value is > than wallet balance, output an error
// GetStateNonce(), GetPrivateKey()
// Convert recipient to []byte; if unsuccessful output an error
// MakeContract(...) (use version global), SignContract(...)
// Output a contract message, with the following structure:
// producer.SecretBytes + uint8(1) + serializedContract
// NOTE: The uint8(1) here will let the producer know that this is a contract message
func ContractMessageFromInput(version uint16, value string, recipient string) (*Contract, error) {
	intVal, err := strconv.Atoi(value) // convert value (string) to int
	if err != nil {
		return nil, errors.New("Unable to convert input to int " + err.Error())
	}

	// case input is zero or less
	if intVal <= 0 {
		return nil, errors.New("Input value is less than or equal to zero")
	}

	// case balance < input
	balance, err := wallet.GetBalance()
	if err != nil {
		return nil, errors.New("Failed to get balance: " + err.Error())
	}
	if balance < uint64(intVal) {
		return nil, errors.New("Input is greater than available balance")
	}

	stateNonce, err := wallet.GetStateNonce()
	if err != nil {
		return nil, errors.New("Failed to get stateNonce: " + err.Error())
	}

	// case recipBytes != 32
	recipBytes, err := hex.DecodeString(recipient)
	if err != nil {
		return nil, errors.New("Failed to hex decode recipient")
	}
	if len(recipBytes) != 32 {
		return nil, errors.New("Failed to convert recipient to size 32 byte slice")
	}

	senderPubKey, err := wallet.GetPrivateKey()
	if err != nil {
		return nil, err
	}

	contract, err := MakeContract(version, senderPubKey, recipBytes, uint64(intVal), stateNonce+1)
	if err != nil {
		return nil, err
	}
	contract.SignContract(senderPubKey)

	return contract, nil
}
