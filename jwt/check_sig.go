package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"strings"
)

// For HMAC signing method, the key can be any []byte. It is recommended to generate
// a key using crypto/rand or something equivalent. You need the same key for signing
// and validating.
var hmacSampleSecret []byte

func main() {
	// Load sample key data
	if keyData, e := ioutil.ReadFile("keys/hmac_key"); e == nil {
		hmacSampleSecret = keyData
	} else {
		panic(e)
	}

	token, err := ExampleNew_hmac()

	if err != nil {
		log.Fatal(err)
	}

	//	ExampleParse_hmac(token)
	fmt.Printf("token: %s\n", token)
	success, message := ParseJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsZXZlbCI6InVzZXIiLCJ1c2VyIjoic2lkIn0.u1v8W3-6vmnpnU1foWRNTChnzuhWRMWjtR5NalDYQKs")

	if success {
		fmt.Printf("Login success\n")
		fmt.Printf("Message: %s\n", message)
	} else {
		fmt.Printf("Login failed\n")
		fmt.Printf("Message: %s\n", message)
	}
	// also works with a d on the end
	// ExampleParse_hmac("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJ4Zm9vIjoiYmFyIn0.uE9pOf9a6USdWHmDx7lxcWrdpjndc0oFUNuVR5GXKLc")
}

// Example creating, signing, and encoding a JWT token using the HMAC signing method
func ExampleNew_hmac() (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":  "sid",
		"level": "user",
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)

	fmt.Printf("The newly created token is %s\n", tokenString)

	return tokenString, err
}

func getToken(token *jwt.Token) (interface{}, error) {
	// Don't forget to validate the alg is what you expect:
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	return hmacSampleSecret, nil
}

// Example parsing and validating a token using the HMAC signing method
func ParseJWT(tokenString string) (bool, string) {
	fmt.Println("In checker")

	var success bool = false
	var message string = ""

	token, err := jwt.Parse(tokenString, getToken)

	if token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			fmt.Printf("The token passed in is: %s\n", tokenString)
			fmt.Printf("Raw: %s\n", token.Raw)
			fmt.Printf("Claims: %u\n", token.Claims)
			payload := strings.Split(token.Raw, ".")[1]
			fmt.Printf("Payload: %s\n", payload)

			fmt.Printf("Claims - Foo: %s\n\n", claims["foo"])
			//fmt.Printf("Foo: %s\nnbf: %f\n", claims["foo"], claims["nbf"])

			message = fmt.Sprintf("Welcome %s (%s)", claims["user"], claims["level"])
			success = true
		}
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		// This is from <https://godoc.org/github.com/dgrijalva/jwt-go#pkg-constants>
		if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
			fmt.Println("Invalid signature")
			fmt.Printf("The signature submitted is: %s\n", token.Signature)

			newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, token.Claims)
			newTokenString, _ := newToken.SignedString(hmacSampleSecret)

			newParsedToken, _ := jwt.Parse(newTokenString, getToken)
			message = fmt.Sprintf("The new signature is: %s\n", newParsedToken.Signature)
		} else if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Printf("There was an error parsing the token: %u", err)
	}

	return success, message
}
