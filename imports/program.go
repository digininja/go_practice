package main

import "fmt"
import "github.com/digininja/go_practice/imports/SystemInfo"

func main() {
	info := NetworkInfoPackage.NetworkInfo()
	fmt.Printf("Returned: %s\n", info)

	// This will fail because function name starts with lower case
	/*
		private := NetworkInfoPackage.notImportable()
		fmt.Printf("Returned: %s\n", private)
	*/
}
