package main

import (
	"fmt"
	"github.com/agnivade/levenshtein"
)

func main() {
	source := "kitten"
	target := "sitting"

	source = "banana"
	target = "apple"

	source = "Orange"
	target = "Apple"
	distance := levenshtein.ComputeDistance(source, target)
	fmt.Printf("The distance between '%s' and '%s' computed as %d.\n", source, target, distance)
	// Output:
	// The distance between kitten and sitting is 3.
}
