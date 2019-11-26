package addressbook

import (
	"fmt"
	"testing"
)

func TestSaveWalletAddress(t *testing.T) {
	// defer os.Remove(constants.AddressBook)

	fmt.Println(SaveWalletAddress("angle", "12345"))
	SaveWalletAddress("blah", "09876")
	// // recipients of different bytes to test
	// tests := []struct {
	// 	name  string
	// 	recip string
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
