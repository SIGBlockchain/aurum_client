package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/SIGBlockchain/aurum_client/internal/privatekey"
)

func TestValidRecipLen(t *testing.T) {
	// recipients of different bytes to test
	tests := []struct {
		name  string
		recip string
		want  bool
	}{
		{
			name:  "valid recipient",
			recip: "2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2deded",
			want:  true,
		},
		{
			name:  "blank recipient",
			recip: "",
			want:  false,
		},
		{
			name:  "one byte recipient",
			recip: "5",
			want:  false,
		},
		{
			name:  "63 byte recipient",
			recip: "2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2dede",
			want:  false,
		},
		{
			name:  "74 byte recipient",
			recip: "2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2deded1a2c3c4e3d",
			want:  false,
		},
		{
			name:  "invalid hex characters",
			recip: "2d2d2@2d2d2d2d2d2d2d2L2d2d2d2d2d2d2dm2d2d2d2d2d2d2d2d2d2d2dededQ",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := ValidRecipLen(tt.recip); result != tt.want {
				t.Errorf("Error: using %s\n", tt.name)
			}
		})
	}
}

// TODO RecoverWallet should be implemented with mocking
func TestRecoverWallet(t *testing.T) {
	// arrange
	// Generate ecdsa key pairs
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	// Encodes private key
	pemEncoded, _ := privatekey.EncodePrivateKey(privateKey)

	// Encodes the pem encoded private key into string and stores it into a jsonStruct
	expected := hex.EncodeToString(pemEncoded)

	// act
	err := RecoverWallet(expected)
	if err != nil {
		t.Error("failed to recover wallet")
	}

	tmpfile, err := os.Open("aurum_wallet.json")
	if err != nil {
		t.Error("Failed to open wallet")
	}
	defer tmpfile.Close()

	// Reads the json file and stores the data into a byte slice
	data, err := ioutil.ReadAll(tmpfile)
	if err != nil {
		t.Error("Failed to read wallet")
	}

	// Json struct for storing the data from the json file
	type jsonStruct struct {
		PrivateKey string
	}

	// Parse the data from the json file into a jsonStruct
	var j Wallet
	err = json.Unmarshal(data, &j)
	if err != nil {
		t.Error("failed to unmarshal")
	}

	actual := j.WalletAddress

	// assert
	if actual != expected {
		t.Error("recovered wallet address is not the same as actual wallet address")
	}

	err = tmpfile.Close() //closes file
	if err != nil {
		t.Errorf("Failed to remove database: %s", err)
	}

	err = os.Remove(tmpfile.Name()) //deletes the file
	if err != nil {
		t.Errorf("Failed to remove database: %s", err)

	}
}
