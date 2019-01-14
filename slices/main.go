package main

import "fmt"

const CAP = 5

func uptoCap(len int) int {
	if len > CAP {
		return CAP
	} else {
		return len
	}
}

func main() {
	var data = make([]int, 10)
	data[0] = 0
	data[1] = 1
	data[2] = 2
	data[3] = 3
	data[4] = 4
	data[5] = 5
	data[6] = 6
	data[7] = 7
	data[8] = 8
	data[9] = 9
	for _, val := range data[0:uptoCap(len(data))] {
		fmt.Printf("%d", val)

	}
}
