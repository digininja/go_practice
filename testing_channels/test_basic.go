package main

import "testing"

/*
var messages chan string
var wg sync.WaitGroup
*/
func TestAdd(t *testing.T) {
	messages = make(chan int)

	input := 10
	output := 11

	wg.Add(1)
	go add_one(&wg, input)

	var msg int
	msg = <-messages

	if msg != output {
		t.Errorf("Double %d == %d, want %d", input, msg, output)
	}

}
func TestDouble(t *testing.T) {
	messages = make(chan int)

	input := 15
	output := 30

	wg.Add(1)
	go double(&wg, input)

	var msg int
	msg = <-messages

	if msg != output {
		t.Errorf("Double %d == %d, want %d", input, msg, output)
	}

}
