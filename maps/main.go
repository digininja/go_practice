package main

import "fmt"
import "math/rand"
import "strconv"

// Intereting guide
// https://blog.golang.org/go-maps-in-action

//map[int]string)
func doit(key string, pos int, amap map[string]map[int]string) string {
	if _, ok := amap[key]; !ok {
		anarr := map[int]string{}
		anarr[pos] = "bbblllaaahhh" + strconv.Itoa(pos)
		amap[key] = anarr
		return amap[key][pos]
	}
	amap[key][pos] = "exitsts" + strconv.Itoa(pos)

	return "exists"

}

func main() {
	fmt.Println("vim-go")
	//mymap := map[string]string{}
	mymap := map[string]map[int]string{}
	//mymap["robin"][0] = "bbb"
	a := doit("robin", 1, mymap)
	fmt.Printf("result: %q\n", a)
	fmt.Printf("result: %q\n", mymap["robin"])

	a = doit("robin", 2, mymap)
	fmt.Printf("result: %q\n", a)
	fmt.Printf("result: %q\n", mymap["robin"])

	a = doit("asdfasdfsdf", 1, mymap)
	fmt.Printf("result: %q\n", a)

	a = doit("wood", 1, mymap)
	fmt.Printf("result: %q\n", a)

	fmt.Printf("map is: %q\n", mymap)

	var data = make(map[string]int)

	for i := 0; i < 10; i++ {
		data[strconv.Itoa(i)] = rand.Intn(50)
	}

	keys := make([]string, len(data))

	i := 0
	for k := range data {
		keys[i] = k
		i++
	}

	//keys := make([]string, 0, len(data))
	fmt.Printf("%q\n", data)
	fmt.Printf("%q\n", keys)
}
