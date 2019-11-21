package addressbook

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	Name          string
	WalletAddress string
}

type AddressBook struct {
	Contacts []Person
}

// SaveWalletAddress adds the wallet address to the address book json file
func SaveWalletAddress(name string, walletAddress string) error {
	addressBook := "address_book.json"
	// check if json exists
	if _, err := os.Stat(addressBook); os.IsNotExist(err) {
		// create json here and forget error if it does not exist?
		return errors.New("wallet file not detected: " + err.Error())
	}
	// open json to write into
	file, err := os.Open(addressBook)
	if err != nil {
		return errors.New("failed to open address book: " + err.Error())
	}
	defer file.Close()

	// Reads the file and stores the data into a byte slice
	jsonData, err := ioutil.ReadAll(file)
	if err != nil {
		return errors.New("Failed to read wallet")
	}

	fmt.Println("jsonData")
	fmt.Println(jsonData)

	// gets the current state of the address book
	var currentBook AddressBook
	err = json.Unmarshal(jsonData, &currentBook)
	if err != nil {
		return err
	}

	// // create new entry
	// var newEntry Person
	// newEntry.Name = name
	// newEntry.WalletAddress = walletAddress

	return nil
}
