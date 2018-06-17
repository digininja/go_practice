package main

import "fmt"
import "math"

type Circle interface {
	Circumference() float64
}

type Tree interface {
	Branches() int
	Leaves() int
}

type MyBigStruct struct {
	length   int
	width    int
	diameter int
	colour   string
	branches int
	leaves   int
}

func (m MyBigStruct) PrintColour() {
	fmt.Printf("Colour: %s\n", m.colour)
}

func (m MyBigStruct) Leaves() int {
	return (m.leaves)
}

func (m MyBigStruct) Branches() int {
	return (m.branches)
}

func (m MyBigStruct) Circumference() float64 {
	return float64(m.diameter) * math.Pi
}

func PrintCircleInfo(c Circle) {
	fmt.Printf("Circ: %f\n", c.Circumference())
}

func PrintTreeStuff(t Tree) {
	fmt.Printf("Tree: %d %d\n", t.Branches(), t.Leaves())
}

func main() {
	var m MyBigStruct

	m.diameter = 40
	m.branches = 30
	m.leaves = 20
	m.colour = "red"

	PrintTreeStuff(m)
	PrintCircleInfo(m)

	m.PrintColour()

	// This won't work
	// fmt.Printf("Circ: %f\n", c.width())
}
