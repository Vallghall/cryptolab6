package main

import (
	"cryptolab6/pkg/rsa"
	"flag"
)

var (
	method = flag.String("m", "RSA", "Choice of a method for an e-signature ciphering")
	text   = flag.String("txt", "ExampleText", "Text for an e-signature ciphering")
	text2  = flag.String("rtxt", "ExampleText", "Testcase text for checking control sums")
)

func main() {
	flag.Parse()
	switch *method {
	case "RSA":
		rsa.DoRSASigning(*text, *text2)
	case "EG":

	}
}
