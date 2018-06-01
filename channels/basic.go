package main

import "sync"
import "fmt"

var messages chan string
var wg sync.WaitGroup

func add_one(wg *sync.WaitGroup, val int) {
	defer wg.Done()
	fmt.Printf("add %d\n", val)
	messages <- fmt.Sprintf("%d", val+1)
}

func double(wg *sync.WaitGroup, val int) {
	defer wg.Done()
	fmt.Printf("double %d\n", val)
	messages <- fmt.Sprintf("%d", val*2)
}

func handle_ret() {
	var msg string
	for {
		msg = <-messages
		if msg == "done" {
			fmt.Printf("closing down handle_ret\n")
			break
		}
		fmt.Printf("Message passed: %s\n", msg)
	}
}

func main() {
	messages = make(chan string)

	go handle_ret()

	for i := 0; i < 10000; i++ {
		wg.Add(2)
		go add_one(&wg, i)
		go double(&wg, i)
	}

	wg.Wait()

	fmt.Printf("Finished\n")
}
