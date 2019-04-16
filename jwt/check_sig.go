package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"time"
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

	ExampleParse_hmac(token)
	BadValidate(token)
}

// Example creating, signing, and encoding a JWT token using the HMAC signing method
func ExampleNew_hmac() (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo":    "bar",
		"nbfx":   time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"expold": time.Now().Add(time.Second * 3600 * 24).Unix(),
		"exp":    time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)

	fmt.Printf("The token is %s.\n", tokenString)

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
func ExampleParse_hmac(tokenString string) {
	fmt.Println("In good checker")

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.

	// getToken is a function, not a variable and not a call to a function

	token, err := jwt.Parse(tokenString, getToken)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Printf("Foo: %s\nnbf: %f\n", claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}

	// Output: bar 1.4444784e+09
}

// Override time value for tests.  Restore default value after.
func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}

func BadValidate(tokenString string) {
	fmt.Println("In bad checker")
	// Override time so we don't care about expired time
	at(time.Unix(0, 0), func() {
		token, err := jwt.Parse(tokenString, getToken)

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Printf("%v %v\n", claims["foo"], claims["exp"])
		} else {
			fmt.Println(err)
		}
	})
}
