package contracts

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/binary"
	"errors"

	"github.com/SIGBlockchain/aurum_client/pkg/hashing"
	"github.com/SIGBlockchain/aurum_client/pkg/publickey"
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
