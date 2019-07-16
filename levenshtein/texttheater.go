package main

import "fmt"
import "github.com/texttheater/golang-levenshtein/levenshtein"

/*
Doesn't work, raised ticket

https://github.com/texttheater/golang-levenshtein/issues/11
*/

func main() {
	source := "kitten"
	target := "sitting"
	distance := levenshtein.DistanceForStrings([]rune(source), []rune(target), levenshtein.DefaultOptions)
	fmt.Printf("The distance between '%s' and '%s' computed as %d.\n", source, target, distance)
}
