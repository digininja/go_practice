package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Configuration struct {
	User      string
	Key       string
	MinLength int
	Users     []string
	Groups    []string
}

func main() {
	configFile := "config.json"
	if _, err := os.Stat(configFile); err != nil {
		fmt.Println("Config file not found")
		return
	}
	file, err := os.Open(configFile)
	if err != nil {
		log.Fatal("Error opening config file")
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(configuration.User)
	if configuration.Key == "" {
		fmt.Println("Key missing from config file")
		return
	}
	fmt.Println(configuration.Key)
	fmt.Println(configuration.MinLength)
	for _, user := range configuration.Users {
		fmt.Println(user)
	}
	fmt.Println(configuration.Groups)
}
