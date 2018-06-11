package main

import "fmt"

func main() {
	fmt.Printf("%f\n", 12)
	fmt.Printf("%.2f\n", float64(12))
	fmt.Printf("%.2f\n", 12.34)
	fmt.Printf("%.2f\n", 12.3)
	fmt.Printf("%.2f\n", 12.3456)
}
