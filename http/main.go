package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getPage(page string) {
	resp, err := http.Get(page)
	if err != nil {
		fmt.Printf("Could not connect: %s\n", err)
	} else {
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading body: %s\n", err)
		} else {
			fmt.Printf("Body: %s\n", string(body))
		}
	}
}
func main() {
	// A basic get
	getPage("https://digi.ninja")

	// It will automatically follow redirects
	getPage("http://digi.ninja")

	// A failure
	getPage("http://digix.ninja")
}
