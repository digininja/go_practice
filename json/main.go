package main

import "fmt"
import "encoding/json"

// The properties have to have upper case letters
// otherwise the Marshal call ignores them and you
// get an empty JSON object
type test struct {
	Abool bool
	Astr  string
}

func main() {
	fmt.Println("vim-go")

	t := test{true, "Hello"}

	b, err := json.Marshal(t)
	if err != nil {
		fmt.Println("error")
	}

	fmt.Printf("Flattened is %s\n", b)

}
