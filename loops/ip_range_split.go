package main

import "fmt"

func main() {
	for i := 0; i < 256; i++ {
		fmt.Printf("echo 10.210.%d.0/24\n", i)
	}
}
