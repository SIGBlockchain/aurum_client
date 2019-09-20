package publickey

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// Returns the PEM-Encoded byte slice from a given public key
func EncodePublicKey(key *ecdsa.PublicKey) []byte {
	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(key)
	return pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})
}

// Decode returns the public key from a given PEM-Encoded byte slice representation of the public key or a non-nil error if fail
func Decode(key []byte) (*ecdsa.PublicKey, error) {
	if key == nil {
		return nil, errors.New("Could not return the decoded public key - the key value is nil")
	}
	blockPub, _ := pem.Decode(key)
	// pem.Decode will return nil for the first value if no PEM data is found. This would be bad
	if blockPub == nil {
		return nil, errors.New("Could not return the public key - failed to PEM decode in preparation x509 encode")
	}

	x509EncodedPub := blockPub.Bytes
	genericPublicKey, err := x509.ParsePKIXPublicKey(x509EncodedPub)
	if err != nil {
		return nil, err
	}
	return genericPublicKey.(*ecdsa.PublicKey), nil
}
