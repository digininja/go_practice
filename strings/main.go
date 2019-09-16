package main

import (
	"fmt"
)

/*

https://blog.golang.org/strings

*/

func main() {
	// Print the character number and then the character for each item in the string

	for i, rune := range "hello world" {
		fmt.Printf("%d: %c\n", i, rune)
	}
}
