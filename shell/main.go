package main

/*
This is a good reference for running commands:

https://nathanleclaire.com/blog/2014/12/29/shelled-out-commands-in-golang/
*/

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "git"
	cmdArgs := []string{"log", "--pretty=full", "--name-only", "--all"}

	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		log.Fatal(fmt.Sprintf("There was an error running git command: %s", err))
	}
	outputStr := string(cmdOut)
	// fmt.Println("Output from command: %s", outputStr)

	scanner := bufio.NewScanner(strings.NewReader(outputStr))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "commit") {
			fmt.Printf("Commit ID: %s\n", line)
		}
		//fmt.Println(scanner.Text())
	}
}
