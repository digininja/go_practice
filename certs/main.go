package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem" // needed for debug writing out csr
	"fmt"
	"io/ioutil"
	"os"
)
import log "github.com/sirupsen/logrus"

// https://golang.org/src/crypto/x509/example_test.go

func LoadX509KeyPair(certFile, keyFile string) (*x509.Certificate, *x509.Certificate, *rsa.PrivateKey) {
	// The file should contain the CA certificate followed by the site certificate
	cf, e := ioutil.ReadFile(certFile)
	if e != nil {
		log.Fatalf("Error loading the certificate file, error:", e.Error())
	}

	kf, e := ioutil.ReadFile(keyFile)
	if e != nil {
		log.Fatalf("Error loading the private key file, error:", e.Error())
	}

	log.Debug("Decoding the first part of the file, should be the CA")
	caBlock, rest := pem.Decode(cf)

	var certBlock *pem.Block
	if rest != nil {
		log.Debug("Decoding the second part of the file, should be the site certificate")
		certBlock, rest = pem.Decode(rest)
	} else {
		log.Fatal("Only one certificate in the file")
	}

	log.Debug("Decoding the private key")
	keyBlock, _ := pem.Decode(kf)

	if keyBlock == nil {
		log.Fatal("Private key came back nil")
	}

	ca, e := x509.ParseCertificate(caBlock.Bytes)
	if e != nil {
		log.Fatal("Error parsing CA: %s", e.Error())
		os.Exit(1)
	}

	cert, e := x509.ParseCertificate(certBlock.Bytes)
	if e != nil {
		log.Fatal("Error parsing certificate: %s", e.Error())
		os.Exit(1)
	}

	key, e := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if e != nil {
		log.Fatal("Error parsing private key: %s", e.Error())
		os.Exit(1)
	}
	return cert, ca, key
}

func main() {
	log.SetLevel(log.DebugLevel)
	ca, cert, key := LoadX509KeyPair("cert.pem", "private.key")

	fmt.Println("Certificate Stuff")
	// https://golang.org/pkg/crypto/x509/#Certificate
	fmt.Printf("CA DNS Names%s\n", ca.DNSNames)
	fmt.Println("Certificate DNS Names")
	for _, name := range cert.DNSNames {
		fmt.Printf("\t%s\n", name)
	}
	fmt.Printf("Not Valid After: %s\n", cert.NotAfter)

	// https://golang.org/pkg/crypto/rsa/#PrivateKey
	fmt.Println("Private Key Stuff")
	fmt.Printf("Exponent:\n")
	fmt.Println(key.D)
}
