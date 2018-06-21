package main

import "fmt"

func main() {
	matches := map[int]string{}
	matches[3] = "fizz"
	matches[5] = "buzz"
	matches[7] = "zing"

	for i := 1; i <= 110; i++ {
		output := ""
		for num, word := range matches {
			if _, ok := matches[num]; ok {
				if i%num == 0 {
					output = fmt.Sprintf("%s%s", output, word)
				}
			}
		}
		if output == "" {
			output = fmt.Sprintf("%d", i)
		}
		fmt.Println(output)
	}
}
