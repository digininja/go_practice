package main

import "fmt"
import "sort"

type Pair struct {
	Key   string
	Value int
}

type PairList struct {
	p         []Pair
	sortByKey bool
	reverse   bool
}

func (p PairList) Len() int      { return len(p.p) }
func (p PairList) Swap(i, j int) { p.p[i], p.p[j] = p.p[j], p.p[i] }
func (p PairList) Less(i, j int) bool {
	if p.sortByKey {
		if p.reverse {
			return p.p[i].Key > p.p[j].Key
		} else {
			return p.p[i].Key < p.p[j].Key
		}
	} else {
		if p.reverse {
			return p.p[i].Value > p.p[j].Value
		} else {
			return p.p[i].Value < p.p[j].Value
		}
	}
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

	var pl = PairList{}
	pl.p = append(pl.p, p1)
	pl.p = append(pl.p, p2)
	pl.p = append(pl.p, p3)
	pl.p = append(pl.p, p4)

	fmt.Printf("Before\n")
	for _, val := range pl.p {
		fmt.Printf("%s: %d\n", val.Key, val.Value)
	}

	pl.sortByKey = false
	sort.Sort(pl)
	fmt.Printf("Sort by value\n")
	for _, val := range pl.p {
		fmt.Printf("%s: %d\n", val.Key, val.Value)
	}

	pl.sortByKey = true
	sort.Sort(pl)
	fmt.Printf("Sort by key\n")
	for _, val := range pl.p {
		fmt.Printf("%s: %d\n", val.Key, val.Value)
	}

	pl.sortByKey = true
	pl.reverse = true
	sort.Sort(pl)
	fmt.Printf("Reverse sort by key\n")
	for _, val := range pl.p {
		fmt.Printf("%s: %d\n", val.Key, val.Value)
	}
}
