package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"crypto/tls"
	"crypto/x509"

	"crypto/rsa"

	"github.com/crewjam/saml/samlsp"
)

func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("call to the hello function")
	fmt.Fprintf(w, "Hello, %s!", samlsp.Token(r.Context()).Attributes.Get("cn"))
}

func main() {
	keyPair, err := tls.LoadX509KeyPair("myservice.cert", "myservice.key")
	if err != nil {
		panic(err) // TODO handle error
	}
	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		panic(err) // TODO handle error
	}

	idpMetadataURL, err := url.Parse("https://samltest.id/saml/idp")
	if err != nil {
		panic(err) // TODO handle error
	}

	rootURL, err := url.Parse("http://localhost:8000")
	if err != nil {
		panic(err) // TODO handle error
	}

	log.Println("Creating the SP")
	samlSP, _ := samlsp.New(samlsp.Options{
		IDPMetadataURL: idpMetadataURL,
		URL:            *rootURL,
		Key:            keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate:    keyPair.Leaf,
	})
	log.Println("Created")
	app := http.HandlerFunc(hello)
	http.Handle("/hello", samlSP.RequireAccount(app))
	http.Handle("/saml/", samlSP)
	log.Println("Starting the listener")
	http.ListenAndServe(":8000", nil)
	log.Printf("started the app")
}
