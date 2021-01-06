package main

import "fmt"
import "sort"

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func main() {
	arr := []int{1, 2, 3, 4}
	perms := permutations(arr)
	// fmt.Println(perms)

	var totals []int
	for _, val := range perms[:] {
		total := (val[0] * 1000) + (val[1] * 100) + (val[2] * 10) + val[3]
		// fmt.Printf("%d\n", val)
		fmt.Printf("%d %d\n", total, total*5)
		totals = append(totals, total*5)
	}
	sort.Ints(totals)
	fmt.Println(totals)

	arr = []int{1, 2, 3, 5}
	perms = permutations(arr)
	// fmt.Println(perms)

	totals = nil
	for _, val := range perms[:] {
		total := (val[0] * 1000) + (val[1] * 100) + (val[2] * 10) + val[3]
		// fmt.Printf("%d\n", val)
		fmt.Printf("%d %d\n", total, total*4)
		totals = append(totals, total*4)
	}
	sort.Ints(totals)
	fmt.Println(totals)

	arr = []int{1, 2, 4, 5}
	perms = permutations(arr)
	// fmt.Println(perms)

	totals = nil
	for _, val := range perms[:] {
		total := (val[0] * 1000) + (val[1] * 100) + (val[2] * 10) + val[3]
		// fmt.Printf("%d\n", val)
		fmt.Printf("%d %d\n", total, total*3)
		totals = append(totals, total*3)
	}
	sort.Ints(totals)
	fmt.Println(totals)

	arr = []int{1, 3, 4, 5}
	perms = permutations(arr)
	// fmt.Println(perms)

	totals = nil
	for _, val := range perms[:] {
		total := (val[0] * 1000) + (val[1] * 100) + (val[2] * 10) + val[3]
		// fmt.Printf("%d\n", val)
		fmt.Printf("%d %d\n", total, total*2)
		totals = append(totals, total*2)
	}
	sort.Ints(totals)
	fmt.Println(totals)

	arr = []int{2, 3, 4, 5}
	perms = permutations(arr)
	// fmt.Println(perms)

	totals = nil
	for _, val := range perms[:] {
		total := (val[0] * 1000) + (val[1] * 100) + (val[2] * 10) + val[3]
		// fmt.Printf("%d\n", val)
		fmt.Printf("%d %d\n", total, total*1)
		totals = append(totals, total*1)
	}
	sort.Ints(totals)
	fmt.Println(totals)
}
