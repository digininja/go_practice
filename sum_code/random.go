package main

import "fmt"
import "time"
import "math/rand"

func main() {
	solve := []int{}
	for i := 1; i < 27; i++ {
		solve = append(solve, i)
	}

	rand.Seed(time.Now().UnixNano())

	for {
		rand.Shuffle(len(solve), func(i, j int) { solve[i], solve[j] = solve[j], solve[i] })
		/*	fmt.Printf("%d\n", solve) */

		if !(solve[0]*solve[1] == solve[2]) {
			continue
		}
		if !((solve[24] + solve[22]) == (solve[17] + solve[25])) {
			continue
		}
		if !(solve[18]+solve[3] == solve[23]) {
			continue
		}
		if !(solve[21]+solve[4] == solve[22]) {
			continue
		}
		if !(solve[19]+solve[18] == solve[20]) {
			continue
		}
		if !(solve[6]*solve[16] == solve[17]) {
			continue
		}
		if !(solve[14]*solve[3] == solve[15]) {
			continue
		}
		if !(solve[0]*solve[3] == solve[12]) {
			continue
		}
		if !(solve[11]*solve[6] == solve[10]) {
			continue
		}
		if !(solve[8]+solve[9] == solve[10]) {
			continue
		}
		if !(solve[7]*solve[7] == solve[8]) {
			continue
		}
		if !(solve[5]*solve[6] == solve[2]) {
			continue
		}
		if !(solve[3]*solve[4] == solve[2]) {
			continue
		}
		fmt.Printf("Solution: %d\n", solve)
	}
}
