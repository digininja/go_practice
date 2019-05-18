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
	fmt.Fprintf(w, "Hello, %s %s!\n", samlsp.Token(r.Context()).Attributes.Get("givenName"), samlsp.Token(r.Context()).Attributes.Get("sn"))
	fmt.Fprintf(w, "Phone number: %s\n", samlsp.Token(r.Context()).Attributes.Get("telephoneNumber"))

	/*
	   {
	     "aud": "http://localhost:9000/saml/metadata",
	     "exp": 1558212437,
	     "iat": 1558208837,
	     "nbf": 1558208837,
	     "sub": "AAdzZWNyZXQxC/qR00pUluWrwB9I0ByFJSeoY06vNwzEUd4QXn9SuF6TwfAxYktTQILzGwfXudgtPVLyQixs29C+IOxv2OoPYHHuf/ZxD7Tw8trBQHAyOHzHua6YaclwQmaC7xNYit351QsQg+Mpq+MM",
	     "attr": {
	       "displayName": [
	         "Sheldor"
	       ],
	       "givenName": [
	         "Sheldon"
	       ],
	       "mail": [
	         "scooper@samltest.id"
	       ],
	       "role": [
	         "employee@samltest.id"
	       ],
	       "sn": [
	         "Cooper"
	       ],
	       "telephoneNumber": [
	         "+1-555-555-5515"
	       ],
	       "uid": [
	         "sheldon"
	       ],
	       "urn:oasis:names:tc:SAML:attribute:subject-id": [
	         "scooper@samltest.id"
	       ]
	     }
	   }
	*/

}

func main() {
	port := 9000

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

	rootURL, err := url.Parse(fmt.Sprintf("http://localhost:%d", port))
	if err != nil {
		panic(err) // TODO handle error
	}

	samlSP, _ := samlsp.New(samlsp.Options{
		IDPMetadataURL: idpMetadataURL,
		URL:            *rootURL,
		Key:            keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate:    keyPair.Leaf,
	})

	log.Println("You can create the manifest file like this")
	log.Printf("curl localhost:%d/saml/metadata > digi.xml", port)
	log.Println("")
	log.Println("And then upload it to here")
	log.Println("https://samltest.id/upload.php")
	log.Println("")

	log.Println("Download their data from")
	log.Println("https://samltest.id/download/")
	log.Println("")

	log.Println("Log file is at")
	log.Println("https://samltest.id/logs/idp.log")
	log.Println("")

	app := http.HandlerFunc(hello)
	http.Handle("/hello", samlSP.RequireAccount(app))
	http.Handle("/saml/", samlSP)
	log.Println("Starting the listener")
	log.Printf("Connect on http://localhost:%d/hello\n", port)

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
