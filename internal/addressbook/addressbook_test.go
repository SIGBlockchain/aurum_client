package addressbook

import (
	"fmt"
	"os"
	"testing"

	"github.com/SIGBlockchain/aurum_client/internal/constants"
)

func TestSaveWalletAddressFirstTime(t *testing.T) {

	var testBook AddressBook
	var person1 Person
	person1.Name = "Kali"
	person1.WalletAddress = "2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2deded"

	SaveWalletAddress("Kali", "2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2deded")
	defer os.Remove(constants.AddressBook)

	address, _ := SearchForContact("Kali")
	fmt.Println("address for Kali: ", address)
}

func TestSaveWalletAddressNotFirstTime(t *testing.T) {
	defer os.Remove(constants.AddressBook)

	SaveWalletAddress("angle", "12345")
	SaveWalletAddress("blah", "09876")

	address, _ := SearchForContact("angle")
	fmt.Println("address for angle: ", address)

	// if address book does not exist create and input the one wallet address

	// if already exist and adding another address

	// // recipients of different bytes to test
	// tests := []struct {
	// 	name  string
	// 	walletAddress string
	// 	want  bool
	// }{
	// 	{
	// 		name:  "valid recipient",
	// 		recip: "2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2deded",
	// 		want:  true,
	// 	},
	// }

	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if result := ValidRecipLen(tt.recip); result != tt.want {
	// 			t.Errorf("Error: using %s\n", tt.name)
	// 		}
	// 	})
	// }
}
