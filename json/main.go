package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// The properties have to have upper case letters
// otherwise the Marshal call ignores them and you
// get an empty JSON object
type test struct {
	Abool bool
	Astr  string
}

type Message struct {
	Name    string
	Body    string
	Time    int64
	Missing string
}

func main() {
	t := test{true, "Hello"}

	b, err := json.Marshal(t)
	if err != nil {
		fmt.Println("error")
	}

	fmt.Printf("Flattened is %s\n", b)

	// This is unmarshalling an object with an array at the top level
	// Some of the things in the object are missing and some of the
	// things in the struct are not in the object, but it all still
	// works

	b = []byte(`[{"Other":"Blah","Name":"Alice","Body":"Hello","Time":1294706395881547000},{"Other":"Blah","Name":"Bob","Body":"Hello","Time":1294706395881547000}]`)

	messages := make([]Message, 0)
	err = json.Unmarshal(b, &messages)

	if err != nil {
		fmt.Printf("Error unmarshalling: %s", err.Error)
		log.Fatal("Bye")
	}
	log.Printf("%v", messages)

	for _, message := range messages {
		log.Printf("Name: %s", message.Name)
	}

}
