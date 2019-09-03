package publickey

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
)

// Returns the PEM-Encoded byte slice from a given public key
func EncodePublicKey(key *ecdsa.PublicKey) []byte {
	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(key)
	return pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})
}
