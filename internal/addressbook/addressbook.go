package addressbook

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/SIGBlockchain/aurum_client/internal/constants"
	"github.com/SIGBlockchain/aurum_client/internal/wallet"
)

type Person struct {
	Name          string
	WalletAddress string
}

type AddressBook struct {
	Contacts []Person
}

// Creates a blank address book for the user
func CreateAddressBook(name string, walletAddress string) error {
	// check if json exists
	_, err := os.Stat(constants.AddressBook)
	if err == nil {
		return errors.New("JSON file for address_book already exists")
	}

	// Create JSON file for address book if there isnt one in existance
	file, err := os.Create(constants.AddressBook)
	if err != nil {
		return err
	}
	defer file.Close()

	firstEntry := Person{Name: name, WalletAddress: walletAddress}
	fmt.Printf("first entry: %s, %s\n", firstEntry.Name, firstEntry.WalletAddress)
	var newAddressBook AddressBook
	newAddressBook.Contacts = append(newAddressBook.Contacts, firstEntry)
	jsonEncoded, err := json.Marshal(newAddressBook)
	if err != nil {
		return err
	}

	_, err = file.Write(jsonEncoded)
	if err != nil {
		return err
	}

	fmt.Println("address book created with first entry")
	return nil
}

// SaveWalletAddress adds the wallet address to the address book json file
func SaveWalletAddress(name string, walletAddress string) error {
	addressBook := "address_book.json"

	//check if wallet address is valid length
	if !wallet.ValidRecipLen(walletAddress) {
		return errors.New("Wallet address is not the correct length")
	}

	// create new entry
	var newEntry Person
	newEntry.Name = name
	newEntry.WalletAddress = walletAddress

	// check if json exists
	if _, err := os.Stat(addressBook); err != nil {
		// if does not exist create it and add in the contact
		err = CreateAddressBook(name, walletAddress)
		if err != nil {
			log.Fatalf("Failed to create address book and enter the first entry: %v\n", err)
		}
		// successful first entry
		return nil
	}

	// json exists so open json to write into
	file, err := os.Open(addressBook)
	if err != nil {
		return errors.New("failed to open address book: " + err.Error())
	}
	defer file.Close()

	// Reads the file and stores the data into a byte slice
	jsonData, err := ioutil.ReadAll(file)
	if err != nil {
		return errors.New("Failed to read address book")
	}

	// gets the current state of the address book
	var currentBook AddressBook
	if err := json.Unmarshal(jsonData, &currentBook); err != nil {
		return errors.New("failed to unmarshall current addressbook: %s" + err.Error())
	}

	fmt.Println("current book", currentBook)

	currentBook.Contacts = append(currentBook.Contacts, newEntry)
	fmt.Println("new book", currentBook)
	jsonEncoded, err := json.Marshal(currentBook)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(addressBook, jsonEncoded, 0644); err != nil {
		return errors.New("failed to write to address book file: " + err.Error())
	}

	return nil
}

func SearchForContact(name string) (string, error) {
	// check if json exists
	if _, err := os.Stat(constants.AddressBook); err != nil {
		return "", errors.New("Address book json not found")
	}

	// json exists so open json to search through
	file, err := os.Open(constants.AddressBook)
	if err != nil {
		return "", errors.New("failed to open address book: " + err.Error())
	}
	defer file.Close()

	// Reads the file and stores the data into a byte slice
	jsonData, err := ioutil.ReadAll(file)
	if err != nil {
		return "", errors.New("Failed to read address book")
	}

	// gets the current state of the address book
	var currentBook AddressBook
	if err := json.Unmarshal(jsonData, &currentBook); err != nil {
		return "", errors.New("failed to unmarshall current addressbook: %s" + err.Error())
	}

	// itterate through contacts to find wanted wallet address
	for _, entries := range currentBook.Contacts {
		if entries.Name == name {
			return entries.WalletAddress, nil
		}
	}

	return "", errors.New("Contact was not found")
}
