package main

import "fmt"

/*
This breaks a bigger slice down into smaller chunks.
*/

func main() {
	var data = make([]int, 13)
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
	data[10] = 10
	data[11] = 11
	data[12] = 12

	var chunks [][]int
	chunkSize := 3
	for len(data) > chunkSize {
		chunks = append(chunks, data[0:chunkSize])
		data = data[chunkSize:len(data)]
	}
	chunks = append(chunks, data)

	for i, chunk := range chunks {
		fmt.Printf("Round %d\n", i)
		for _, val := range chunk {
			fmt.Printf("%d\n", val)
		}
	}
}
