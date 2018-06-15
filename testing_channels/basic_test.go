package main

import "testing"

/*
var messages chan string
var wg sync.WaitGroup
*/
func TestChan(t *testing.T) {
	messages = make(chan int)

	input := 1
	output := 2

	wg.Add(1)
	go add_one(&wg, input)

	var msg int
	msg = <-messages

	if msg != output {
		t.Errorf("Double %d == %d, want %d", input, msg, output)
	}

}
