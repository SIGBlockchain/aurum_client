package main

import "github.com/pborman/getopt"

type Opts struct {
	help       *bool
	version    *bool
	setup      *bool
	info       *bool
	updateInfo *bool
	contract   *bool
	recipient  *string
	value      *string
	producer   *string
}

func main() {
	options := Opts{
		help:       getopt.BoolLong("help", '?', "help"),
		version:    getopt.BoolLong("version", 'w', "version"),
		setup:      getopt.BoolLong("setup", 's', "set up client"),
		info:       getopt.BoolLong("info", 'i', "wallet info"),
		updateInfo: getopt.BoolLong("update", 'u', "update wallet info"),
		contract:   getopt.BoolLong("contract", 'c', "make contract"),
		recipient:  getopt.StringLong("recipient", 'r', "recipient"),
		value:      getopt.StringLong("value", 'v', "", "value to send"),
		producer:   getopt.StringLong("producer", 'p', "", "producer address"),
	}

}
