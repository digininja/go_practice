package main

import (
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
)

// https://golang.org/src/crypto/x509/example_test.go

func main() {
	fmt.Println("vim-go")

	asn1Data, err := ioutil.ReadFile("cert.pem") // asn1Data has type []byte
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Read the following: %s", string(asn1Data))
	certs, err := x509.ParseCertificates(asn1Data)
	if err != nil {
		log.Fatal("Could not parse the certificates, error: %s", err)
	}

	log.Printf("certs: %u", certs)

}
