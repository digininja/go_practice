package main

import "fmt"
import "time"
import "math/rand"

func main() {

	var days = []string{
		"monday",
		"tuesday",
		"wednesday",
		"thursday",
		"friday",
		"saturday",
		"sunday",
	}

	var data = make(map[string]int)

	// not crypto secure generateion of seed
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	for _, day := range days {
		data[day] = random.Intn(50)
	}

	for _, k := range days {
		fmt.Println("Day:", k, "Value:", data[k])
	}
}
