package main

import "fmt"
import "github.com/digininja/go_practice/imports/my_module"

func main() {
	data := module.ExposedFunction()
	fmt.Printf("Returned: %s\n", data)

	// This will fail because function name starts with lower case
	/*
		data := module.notExposedFunction()
		fmt.Printf("Returned: %s\n", data)
	*/

	// Because this is in the same module, the function
	// doesn't need to start with a capital
	data = functionInSameModule()
	fmt.Printf("Returned: %s\n", data)
}
