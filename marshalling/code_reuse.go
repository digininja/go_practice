package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

import "encoding/json"

type JSONMessageMarker interface {
	//	foo()
}

type JSONMessage struct {
}

func (r JSONMessage) foo() {
}

type Request1 struct {
	JSONMessage
	Name string
}

type Request2 struct {
	JSONMessage
	Number int
}

func MarshallJSONMessage(msg JSONMessageMarker) (string, error) {
	js, err := json.Marshal(msg)
	if err != nil {
		return "", fmt.Errorf("Error marshalling the JSON request: %s", err.Error())
	}
	s := string(js[:])
	return s, nil
}

func main() {
	req1 := Request1{Name: "robin"}
	s1, _ := MarshallJSONMessage(req1)
	fmt.Printf("Marshalled: %s\n", s1)
	req2 := Request2{Number: 123}
	s2, _ := MarshallJSONMessage(req2)
	fmt.Printf("Marshalled: %s\n", s2)
	s3, _ := MarshallJSONMessage("hello")
	fmt.Printf("Marshalled: %s\n", s3)
	var slice []int
	slice = append(slice, 2)
	slice = append(slice, 1)
	slice = append(slice, 3)
	s4, _ := MarshallJSONMessage(slice)
	fmt.Printf("Marshalled: %s\n", s4)

	var mapData = make(map[string]int)

	for i := 0; i < 10; i++ {
		mapData[strconv.Itoa(i)] = rand.Intn(50)
	}
	s5, _ := MarshallJSONMessage(mapData)
	fmt.Printf("Marshalled: %s\n", s5)

}
