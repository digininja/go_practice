package main

import "fmt"
import "log"
import "encoding/json"

type JSONMessage struct {
}

func (r JSONMessage) Marshall() string {
	js, err := json.Marshal(r)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error marshalling the JSON request: %s", err.Error()))
	}
	s := string(js[:])
	return s
}

type Request1 struct {
	JSONMessage
	Name string
}

func (r Request2) Marshall() string {
	js, err := json.Marshal(r)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error marshalling the JSON request: %s", err.Error()))
	}
	s := string(js[:])
	return s
}

type Request2 struct {
	JSONMessage
	Number int
}

func main() {
	fmt.Println("vim-go")
	req1 := Request1{Name: "robin"}
	fmt.Printf("Marshalled: %s\n", req1.Marshall())
	req2 := Request2{Number: 123}
	fmt.Printf("Marshalled: %s\n", req2.Marshall())
}
