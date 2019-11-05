package wallet

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/SIGBlockchain/aurum_client/internal/constants"
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

func TestContractHistoryExists(t *testing.T) {
	// create json wallet
	SetupWallet()
	defer os.Remove(constants.Wallet)

	// open json file
	jsonFile, err := os.Open(constants.Wallet)
	if err != nil {
		t.Errorf("Error opening JSON")
	}
	defer jsonFile.Close()

	// read file into byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// create wallet struct
	var testWallet Wallet
	json.Unmarshal(byteValue, &testWallet)

	// check if contract history field exists in struct
	if testWallet.ContractHistory != nil {
		t.Errorf("ContractHistory not nil")
	}

}
