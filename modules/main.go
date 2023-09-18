package main

import (
	"fmt"

	/*
		The first word is the name you want to call the package in this script. It does not have to match the `package` statement at the top of the file.
		The second is the  full path to the directory.
	*/
	coreModTwo "github.com/digininja/go_practice/modules/submod2"
	core "github.com/digininja/go_practice/modules/submodules"
)

func main() {
	fmt.Println("Main file.")
	core.PrintStuff()
	coreModTwo.PrintMoreStuff()
	coreModTwo.PrintMoreStuff2()
}
