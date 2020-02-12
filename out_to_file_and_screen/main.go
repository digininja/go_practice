package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	stdout := os.Stdout
	file, err := os.Create("/tmp/dat2")
	check(err)

	defer stdout.Close()
	defer file.Close()

	d1 := []byte{115, 111, 109, 101, 10}
	n1, err := stdout.Write(d1)
	check(err)
	fmt.Printf("wrote %d bytes to standard out\n", n1)
	n2, err := file.Write(d1)
	check(err)
	fmt.Printf("wrote %d bytes to a file\n", n2)

	n3, err := stdout.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes to standard out\n", n3)
	n4, err := file.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes to a file\n", n4)

	stdout.Sync()
	file.Sync()
}
