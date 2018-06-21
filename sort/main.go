package main

import "fmt"
import "sort"

type Pair struct {
	Key   string
	Value int
}

type PairListValueSort []Pair
type PairListKeySort []Pair
type PairList []Pair

func (p PairListValueSort) Len() int      { return len(p) }
func (p PairListValueSort) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairListValueSort) Less(i, j int) bool {
	return p[i].Value < p[j].Value
}

func (p PairListKeySort) Len() int      { return len(p) }
func (p PairListKeySort) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// If keys match then sort by values to keep things in a nice order
func (p PairListKeySort) Less(i, j int) bool {
	if p[i].Key == p[j].Key {
		return (p[i].Value < p[j].Value)
	}
	return p[i].Key < p[j].Key
}

func main() {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted

	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)

	a := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Sort(sort.IntSlice(a))
	fmt.Println(a)

	p1 := Pair{"abc", 123}
	p2 := Pair{"abd", 124}
	p3 := Pair{"abc", 1}
	p4 := Pair{"xyz", 567}
	p5 := Pair{"raw", 2}
	p6 := Pair{"abc", 12}

	var pl = PairList{}
	pl = append(pl, p1)
	pl = append(pl, p2)
	pl = append(pl, p3)
	pl = append(pl, p4)
	pl = append(pl, p5)
	pl = append(pl, p6)

	fmt.Printf("Before\n")
	for _, val := range pl {
		fmt.Printf("%s: %d\n", val.Key, val.Value)
	}

	sort.Sort(PairListValueSort(pl))
	fmt.Printf("Sort by value\n")
	for _, val := range pl {
		fmt.Printf("%s: %d\n", val.Key, val.Value)
	}

	sort.Sort(PairListKeySort(pl))
	fmt.Printf("Sort by key\n")
	for _, val := range pl {
		fmt.Printf("%s: %d\n", val.Key, val.Value)
	}

	sort.Sort(sort.Reverse(PairListKeySort(pl)))
	fmt.Printf("Reverse sort by key\n")
	for _, val := range pl {
		fmt.Printf("%s: %d\n", val.Key, val.Value)
	}
}
